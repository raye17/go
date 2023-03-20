package main

import (
	"k8sClient-go/client"
	"k8sClient-go/secret"
)

func main() {
	//client.Rest()
	_ = client.Clientset()
	//pod.Namespace = "default"
	//pod.CreatePod()
	//pod.ListPod()
	//pod.WatchPod()
	//deployment.CreateDeployment()
	//informer.CreatePodInformer()
	//informer.CreateServiceInformer()
	secret.Secret()

}
