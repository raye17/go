package pod

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sClient-go/client"
	"log"
	"time"
)

var Namespace string

func CreatePod() {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: Namespace,
			Labels: map[string]string{
				"run": "nginx",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Image: "nginx",
					Name:  "nginx-container",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}
	_, err := client.Clientset().CoreV1().Pods(Namespace).Create(
		context.Background(),
		&pod,
		metav1.CreateOptions{})
	if err != nil {
		log.Printf("create pod err,error:%v\n", err)
		return
	}
	log.Println("create pod success!")
}
func ListPod() {
	podList, err := client.Clientset().CoreV1().Pods(corev1.NamespaceAll).List(
		context.Background(),
		metav1.ListOptions{})
	if err != nil {
		log.Printf("list pod error:%v", err)
		return
	}
	fmt.Println("test pod count:", len(podList.Items))
	for _, pod := range podList.Items {
		fmt.Printf("namespace:%s\tpod name:%s\n", pod.Namespace, pod.Name)
	}
}
func WatchPod() {
	watch, err := client.Clientset().CoreV1().Pods(corev1.NamespaceDefault).Watch(context.Background(),
		metav1.ListOptions{})
	if err != nil {
		log.Panicln("watch err:", err)
	}
	defer watch.Stop()
	for {
		item := <-watch.ResultChan()
		pod := item.Object.(*corev1.Pod)
		fmt.Printf("event type:%s pod name:%s %s\n",
			item.Type,
			pod.Name,
			time.Now().Format("2006-01-02 15:04:05"))
	}
}
