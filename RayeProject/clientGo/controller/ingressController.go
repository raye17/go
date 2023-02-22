package controller

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	apinetv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	v1 "k8s.io/client-go/informers/core/v1"
	networkv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	listercorev1 "k8s.io/client-go/listers/core/v1"
	listernetv1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"reflect"
	"sync"
)

type Config struct {
	workerNum int
	maxRetry  int //最大重试次数
}
type IngressController struct {
	client          kubernetes.Interface
	serviceInformer v1.ServiceInformer
	ingressInformer networkv1.IngressInformer
	serviceLister   listercorev1.ServiceLister
	ingressLister   listernetv1.IngressLister
	workQue         workqueue.RateLimitingInterface
	Config
	wg sync.WaitGroup
}

var defaultConfig = Config{5, 8}

type OptionsFunc func(c *Config)

func WithWorkerNum(num int) OptionsFunc {
	return func(c *Config) {
		c.workerNum = num
	}
}
func WithMaxRetry(maxRetry int) OptionsFunc {
	return func(c *Config) {
		c.maxRetry = maxRetry
	}
}
func NewIngressController(client kubernetes.Interface, serviceInformer v1.ServiceInformer,
	ingressInformer networkv1.IngressInformer, opts ...OptionsFunc) *IngressController {
	c := &IngressController{
		client:          client,
		serviceInformer: serviceInformer,
		ingressInformer: ingressInformer,
		workQue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
		serviceLister:   serviceInformer.Lister(),
		ingressLister:   ingressInformer.Lister(),
	}
	c.serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
		DeleteFunc: c.deleteService,
	})
	c.ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})
	c.Config = defaultConfig
	for _, fn := range opts {
		fn(&c.Config)
	}
	return c
}
func (c *IngressController) ShutDown() {
	c.workQue.ShutDown()
}
func (c *IngressController) Run() {
	for i := 0; i < c.Config.workerNum; i++ {
		c.wg.Add(1)
		go c.worker()
	}
	c.wg.Wait()
}
func (c *IngressController) worker() {
	defer c.wg.Done()
	for c.processNextItem() {
	}
}
func (c *IngressController) processNextItem() bool {
	item, shutdown := c.workQue.Get()
	if shutdown {
		return false
	}
	defer c.workQue.Done(item)
	key := item.(string)
	err := c.syncService(key)
	if err != nil {
		c.handleError(key, err)
	}
	return true
}
func (c *IngressController) syncService(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		klog.Errorf("split meta namespace key error:%v\n", err)
		return err
	}
	service, err := c.serviceLister.Services(namespace).Get(name)
	if errors.IsNotFound(err) {
		ingress, err := c.ingressLister.Ingresses(namespace).Get(name)
		if err != nil {
			if errors.IsNotFound(err) {
				return nil
			}
			return err
		}
		ownerReference := metav1.GetControllerOf(ingress)
		if ownerReference == nil {
			return nil
		}
		if ownerReference.Kind == "Service" && ownerReference.Name == name {
			err = c.client.NetworkingV1().Ingresses(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
			if err != nil {
				return err
			}
			klog.Infof("delete ingress,name is %s,namespace is %s\n", name, namespace)
			return nil
		}
		return nil
	}
	if err != nil {
		klog.Errorf("get service error,name: %s,namespace:s%,error:%s\n", name, namespace, err)
		return err
	}
	_, ok := service.GetAnnotations()["ingress/http"]
	ingress, err := c.ingressLister.Ingresses(namespace).Get(name)
	if err != nil && !errors.IsNotFound(err) {
		klog.Errorf("get ingress error,name: %s,namespace:s%,error:%s\n", name, namespace, err)
		return err
	}
	if ok && errors.IsNotFound(err) {
		klog.Infof("create ingress ,name: %s,namespace:%s\n", name, namespace)
		err := c.createIngress(service)
		if err != nil {
			klog.Errorf("create ingress error, name:%s, namespace:%s, err:%v", name, namespace, err)
			return err
		}
	} else if !ok && ingress != nil {
		klog.Infof("delete ingress, name:%s, namespace:%s", name, namespace)
		// 删除ingress
		err := c.client.NetworkingV1().Ingresses(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
		if err != nil {
			klog.Errorf("delete ingress error, name:%s, namespace:%s, err:%v", name, namespace, err)
			return err
		}
	}
	return nil
}
func (c *IngressController) handleError(key string, err error) {
	if c.workQue.NumRequeues(key) <= c.maxRetry {
		c.workQue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.workQue.Forget(key)
}

func (c *IngressController) addService(obj interface{}) {
	c.enqueue(obj)
}
func (c *IngressController) updateService(oldObj interface{}, newObj interface{}) {
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}
func (c *IngressController) deleteService(obj interface{}) {
	c.enqueue(obj)
}
func (c *IngressController) deleteIngress(obj interface{}) {
	ingress := obj.(*apinetv1.Ingress)
	ownerReference := metav1.GetControllerOf(ingress)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind != "Service" {
		return
	}
	c.workQue.Add(ingress.Namespace + "/" + ingress.Name)
}
func (c *IngressController) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.workQue.Add(key)
}
func (c *IngressController) createIngress(service *corev1.Service) error {
	ingress := apinetv1.Ingress{}
	// 将ingress关联到当前service
	ingress.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
		*metav1.NewControllerRef(service, metav1.SchemeGroupVersion.WithKind("Service")),
	}
	ingress.Name = service.Name
	ingress.Namespace = service.Namespace

	pathType := apinetv1.PathTypePrefix
	icn := "nginx"
	ingress.Spec = apinetv1.IngressSpec{
		IngressClassName: &icn,
		Rules: []apinetv1.IngressRule{
			{
				Host: "example.com",
				IngressRuleValue: apinetv1.IngressRuleValue{
					HTTP: &apinetv1.HTTPIngressRuleValue{
						Paths: []apinetv1.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &pathType,
								Backend: apinetv1.IngressBackend{
									Service: &apinetv1.IngressServiceBackend{
										Name: service.Name,
										Port: apinetv1.ServiceBackendPort{
											Number: 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := c.client.NetworkingV1().Ingresses(service.Namespace).Create(context.Background(), &ingress, metav1.CreateOptions{})

	return err
}
