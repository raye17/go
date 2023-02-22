package util

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

const PasswordSecretKey = "userPassword"

func GetPassword(ctx context.Context, secretName string, secrets corev1.SecretInterface) (string, error) {
	var s *v1.Secret
	var err error
	s, err = secrets.Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return "", fmt.Errorf("secret %s not found", secretName)
		}
		return "", fmt.Errorf("failed to get secret %s:%v", secretName, err)
	}
	if password, ok := s.Data[PasswordSecretKey]; ok {
		fmt.Println(string(password))
		return string(password), nil
	}
	return "", errors.New("password not found in secret")

}
