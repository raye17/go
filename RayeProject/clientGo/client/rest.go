package client

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func Rest() {
	// 1构造访问config得的配置，从文件中加载
	config, err := clientcmd.BuildConfigFromFlags("", "conf/config")
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.APIPath = "/api"
	//2 创建rest client
	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// 3查找命名空间下的pod
	var podList v1.PodList
	err = client.Get().Namespace("default").Resource("pods").Do(context.Background()).Into(&podList)
	if err != nil {
		log.Printf("get pods error:%v", err)
		return
	}
	fmt.Println("default pod count:", len(podList.Items))
	for _, pod := range podList.Items {
		fmt.Printf("name: %s\n", pod.Name)
	}
}
