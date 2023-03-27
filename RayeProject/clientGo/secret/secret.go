package secret

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sClient-go/client"
)

type data struct {
	name         string
	Password     string `json:"password"`
	Username     string `json:"Username"`
	UserPassword string `json:"userPassword"`
	Age          string
}

func Secret() {
	data := data{
		name:         "user001",
		Password:     "123456",
		Username:     "username001",
		UserPassword: "123456",
		Age:          "19",
	}
	secret, err := createSecret("test001", "sss", context.TODO(), data)
	if err != nil {
		fmt.Println(err)
	}
	getSecret(secret.Name, secret.Namespace)

}
func createSecret(name string, ns string, ctx context.Context, data data) (*v1.Secret, error) {
	clientSet := client.Clientset()
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Data: map[string][]byte{
			"name":     []byte(data.name),
			"password": []byte(data.Password),
		},
		StringData: map[string]string{
			"Username":     data.Username,
			"userPassword": data.UserPassword,
			"Age":          data.Age,
		},
	}
	_, err := clientSet.CoreV1().Namespaces().Get(ctx, ns, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			names := &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: ns,
				},
			}
			_, err := clientSet.CoreV1().Namespaces().Create(ctx, names, metav1.CreateOptions{})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	Secret, err := clientSet.CoreV1().Secrets(ns).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			Secret, err = clientSet.CoreV1().Secrets(ns).Update(ctx, secret, metav1.UpdateOptions{})
			if err != nil {
				glog.Errorf("failed to update secret:", err)
				return nil, err
			}
		} else {
			glog.Errorf("failed to create secret:", err)
			return nil, err
		}
	}
	return Secret, nil
}
func getSecret(name string, ns string) {
	clientset := client.Clientset()
	secret, err := clientset.CoreV1().Secrets(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	//var info data
	for k, v := range secret.Data {
		fmt.Println(k, ":", string(v))
	}
	//info.name = string(secret.Data["name"])
	//info.Password = string(secret.Data["password"])
	//info.Username = string(secret.Data["Username"])
	//info.UserPassword = string(secret.Data["userPassword"])
	//info.Age = string(secret.Data["Age"])
	//fmt.Printf("%+v", info)
}
