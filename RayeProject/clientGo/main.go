package main

import (
	"fmt"
	"k8sClient-go/client"
	"os/exec"
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
	//secret.Secret()
	config := exec.Command("kubectl", "config", "get-clusters")
	output, _ := config.CombinedOutput()
	s := fmt.Sprintf(string(output[5:]))
	fmt.Println(s)
}
