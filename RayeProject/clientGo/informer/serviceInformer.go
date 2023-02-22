package informer

import (
	"flag"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"k8sClient-go/controller"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func CreateServiceInformer() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 创建config
	var (
		config *rest.Config
		err    error
	)
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Infof("get kubeconfig error, err:%v", err)
		// 如果我们的程序运行在K8S的Pod中，那么就需要下面的方式来获取config
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.Errorf("get incluster config error:%v", err)
			panic(err)
		}
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//clientSet := client.Clientset()
	factory := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()
	ctrl := controller.NewIngressController(clientSet, serviceInformer, ingressInformer, controller.WithWorkerNum(8))
	stopCh := make(chan struct{})
	dealSignal(stopCh, ctrl)
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	klog.Infoln("start success!")
	ctrl.Run()
}
func dealSignal(stopCh chan struct{}, ctrl *controller.IngressController) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh
		klog.Info("program shutdown...")
		close(stopCh)
		ctrl.ShutDown()
	}()
}
