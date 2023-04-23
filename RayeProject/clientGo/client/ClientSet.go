package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Clientset() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", "C:/Users/sunxiaoyang01/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
