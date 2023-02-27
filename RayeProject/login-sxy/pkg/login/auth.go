package login

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	clientset "git.inspur.com/szsciit/cnos/adapter/generated/cnos/clientset/versioned"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"strings"
	"time"
)

const jwtSecret = "login-demo"

var kubeClientSet kubernetes.Interface

type Authentication struct {
	kubeClientSet kubernetes.Interface
	userClientSet clientset.Interface
}

func NewAuthentication(kubeClientSet kubernetes.Interface, userClientSet clientset.Interface) *Authentication {
	return &Authentication{
		kubeClientSet: kubeClientSet,
		userClientSet: userClientSet,
	}
}
func SetKubeClientSet(clientSet kubernetes.Interface) {
	kubeClientSet = clientSet
}

// AuthenticateUser accept a user's credentials and returns a token if valid
func AuthenticateUser(username string, password string) (string, error) {
	//fetch the user object by name
	userSecret, err := kubeClientSet.CoreV1().Secrets("default").Get(context.TODO(), username,
		metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("user %v not found", username)
		}
		return "", fmt.Errorf("failed to get user %s:%v", username, err)
	}
	glog.Infof("get user from secret success")
	hashedPassword, err := base64.StdEncoding.DecodeString(string(userSecret.Data["userPassword"]))
	if err != nil {
		return "", fmt.Errorf("failed decoding user password for user %s:%v", username, err)
	}
	//Verify the provided password against the stored hash
	if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return "", fmt.Errorf("invaild password for user %s:%v", username, err)
	}
	//Generate and return a token for the authenticated user
	token, err := generateToken(username)
	if err != nil {
		return "", fmt.Errorf("failed  generate token for user %s:%v", username, err)
	}
	return token, nil
}

// Middleware to verify user token before allowing access to protected resources
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		//Ensure that the Authorization header is in the correct format and extract the token
		splitToken := strings.Split(authHeader, "Bearer")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		token := splitToken[1]
		username, err := verifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		//Add the authenticated username to the request context
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed generate JWTtoken:%v", err)
	}
	return signedToken, nil
}
func verifyToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Subject, nil
	}
	return "", errors.New("invalid token")
}
