package informer

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

func CreatePodInformer(client *kubernetes.Clientset) {
	clientSet := client
	sharedInformerFactory := informers.NewSharedInformerFactory(clientSet, 0)
	podInformer := sharedInformerFactory.Core().V1().Pods()
	informer := podInformer.Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("[add event] pod name:%s  ns:%s  %s\n",
				pod.Name,
				pod.Namespace,
				time.Now().Format("2006-01-02 15:04:05"))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			pod := newObj.(*v1.Pod)
			fmt.Printf("[update event] pod name:%s  ns:%s  %s\n",
				pod.Name,
				pod.Namespace,
				time.Now().Format("2006-01-02 15:04:05"))
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("[delete pod] pod name:%s  ns:%s  %s\n",
				pod.Name,
				pod.Namespace,
				time.Now().Format("2006-01-02 15:04:05"))
		},
	})
	stopCh := make(chan struct{})
	sharedInformerFactory.Start(stopCh)
	<-stopCh
}
