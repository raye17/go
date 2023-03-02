package main

import (
	"fmt"
	"github.com/golang/glog"
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
	configContext := exec.Command("kubectl", "config", "get-clusters")
	output, _ := configContext.CombinedOutput()
	s := fmt.Sprintf(string(output[5:]))
	fmt.Println(s)
	cluster := fmt.Sprintf("--cluster=%s", s)
	users := fmt.Sprintf("--user=%s", "raye")
	userCmd := exec.Command("kubectl", "config", "set-context", "raye", cluster, users)
	_, err := userCmd.CombinedOutput()
	if err != nil {
		glog.Errorf("failed to set context for user ,error:%v", err)
	}
	glog.Infof("create k8sUser success,and output is %s", s)
}
