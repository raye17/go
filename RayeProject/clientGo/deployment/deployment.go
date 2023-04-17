package deployment

import (
	"bufio"
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsresv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
	"os"
)

func CreateDeployment(client *kubernetes.Clientset) {
	deployClient := client.AppsV1().Deployments("test")
	createDeploy(deployClient)
	prompt()
	updateDeploy(deployClient)
	prompt()
	listDeploy(deployClient)
	prompt()
	deleteDeploy(deployClient)
}
func createDeploy(client appsresv1.DeploymentInterface) {
	klog.Info("createDeploy......")
	replicas := int32(2)
	deploy := appsv1.Deployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "deployment-nginx-demo",
			Namespace: "test",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []corev1.ContainerPort{
								{
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	dep, err := client.Create(context.Background(), &deploy, v1.CreateOptions{})
	if err != nil {
		klog.Errorf("create deployment error:%v\n", err)
		return
	}
	klog.Info("create deployment success,name:", dep.Name)
}
func updateDeploy(client appsresv1.DeploymentInterface) {
	klog.Info("UpdateDeploy...........")
	// 当有多个客户端对同一个资源进行操作时，可能会发生错误。使用RetryOnConflict来重试，重试相关参数由DefaultRetry来提供
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// 查询要更新的deploy
		deploy, err := client.Get(context.Background(), "deployment-nginx-demo", v1.GetOptions{})
		if err != nil {
			klog.Errorf("can't get deployment, err:%v", err)
			return nil
		}

		// 修改参数后进行更新
		replicas := int32(1)
		deploy.Spec.Replicas = &replicas
		deploy.Spec.Template.Spec.Containers[0].Image = "nginx:1.13"

		_, err = client.Update(context.Background(), deploy, v1.UpdateOptions{})
		if err != nil {
			klog.Errorf("update deployment error, err:%v", err)
		}
		return err
	})

	if err != nil {
		klog.Errorf("update deployment error, err:%v", err)
	} else {
		klog.Infoln("update deployment success")
	}
}

func listDeploy(client appsresv1.DeploymentInterface) {
	klog.Info("ListDeploy...........")
	depList, err := client.List(context.Background(), v1.ListOptions{})
	if err != nil {
		klog.Errorf("list deployment error, err:%v", err)
		return
	}

	for _, dep := range depList.Items {
		klog.Infof("deploy name:%s, replicas:%d, container image:%s\n", dep.Name, *dep.Spec.Replicas, dep.Spec.Template.Spec.Containers[0].Image)
	}
}

func deleteDeploy(client appsresv1.DeploymentInterface) {
	klog.Info("DeleteDeploy...........")
	// 删除策略
	deletePolicy := v1.DeletePropagationForeground
	err := client.Delete(context.Background(), "deployment-nginx-demo", v1.DeleteOptions{PropagationPolicy: &deletePolicy})
	if err != nil {
		klog.Errorf("delete deployment error, err:%v", err)
	} else {
		klog.Info("delete deployment success")
	}
}

func prompt() {
	fmt.Printf("-> Press Enter key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
