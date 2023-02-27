package main

import "k8sClient-go/informer"

func main() {
	//client.Rest()
	//client.Clientset()
	//pod.Namespace = "default"
	//pod.CreatePod()
	//pod.ListPod()
	//pod.WatchPod()
	//deployment.CreateDeployment()
	informer.CreatePodInformer()
	//informer.CreateServiceInformer()
}
