package main

import (
	"k8s.io/client-go/kubernetes"
	"k8sClient-go/client"
	"k8sClient-go/pod"
)

var clientSet *kubernetes.Clientset

func init() {
	clientSet = client.Clientset()
}
func main() {
	//client.Rest()
	//pod.Namespace = "default"
	//pod.CreatePod()
	pod.ListPod(clientSet)
	//pod.WatchPod()
	//deployment.CreateDeployment()
	//informer.CreatePodInformer()
	//informer.CreateServiceInformer()
	//secret.Secret()

}
