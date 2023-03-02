/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*package controllers

import (
   "context"
   "k8s.io/client-go/kubernetes"
   "k8s.io/client-go/tools/record"

   "k8s.io/apimachinery/pkg/runtime"
   ctrl "sigs.k8s.io/controller-runtime"
   "sigs.k8s.io/controller-runtime/pkg/client"
   "sigs.k8s.io/controller-runtime/pkg/log"

   cnosv1 "git.inspur.com/szsciit/cnos/adapter/apis/cnos/v1"
)

// UserReconciler reconciles a User object
type UserReconciler struct {
   client.Client
   Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cnos.inspur.com,resources=users,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cnos.inspur.com,resources=users/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cnos.inspur.com,resources=users/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the User object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *UserReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
   _ = log.FromContext(ctx)

   // TODO(user): your logic here

   return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UserReconciler) SetupWithManager(mgr ctrl.Manager) error {
   return ctrl.NewControllerManagedBy(mgr).
      For(&cnosv1.User{}).
      Complete(r)
}
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	apisUerV1 "git.inspur.com/szsciit/cnos/adapter/apis/cnos/v1"
	clientset "git.inspur.com/szsciit/cnos/adapter/generated/cnos/clientset/versioned"
	"git.inspur.com/szsciit/cnos/adapter/generated/cnos/clientset/versioned/scheme"
	cnos "git.inspur.com/szsciit/cnos/adapter/generated/cnos/informers/externalversions"
	listers "git.inspur.com/szsciit/cnos/adapter/generated/cnos/listers/cnos/v1"
	"git.inspur.com/szsciit/cnos/adapter/pkg/util"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	corev1 "k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	controllerName = "user controller"
	ClusterRole    = "cluster-admin"
	CACRT          = "/etc/kubernetes/pki/ca.crt"
	CAKEY          = "/etc/kubernetes/pki/ca.key"
	DAY            = "365"
	AdminUser      = "AdminRole"
	NormalUser     = "NormalRole"
	UserKeyFile    = "key"
	UserCsrFile    = "csr"
	UserCrtFile    = "crt"
)

type UserController struct {
	kubeClientSet kubernetes.Interface
	userClientSet clientset.Interface
	userFactory   cnos.SharedInformerFactory
	userLister    listers.UserLister
	userSynced    cache.InformerSynced
	recorder      record.EventRecorder
}

func NewUserController(kubeClientSet kubernetes.Interface, userClientSet clientset.Interface,
	userFactory cnos.SharedInformerFactory) *UserController {
	//TODO: 写controller逻辑
	userInformer := userFactory.Cnos().V1().Users()
	informer := userInformer.Informer()
	lister := userInformer.Lister()
	runtime.Must(scheme.AddToScheme(scheme.Scheme))
	runtime.Must(apisUerV1.AddToScheme(scheme.Scheme))
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClientSet.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerName})
	controller := &UserController{
		kubeClientSet: kubeClientSet,
		userClientSet: userClientSet,
		userFactory:   userFactory,
		userLister:    lister,
		userSynced:    informer.HasSynced,
		recorder:      recorder,
	}
	glog.Info("setting up user event handles...")
	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    controller.addUser,
		UpdateFunc: controller.updateUser,
		DeleteFunc: func(obj interface{}) {
			user := obj.(*apisUerV1.User)
			glog.Infof("user %s is deleted", user.Name)
		},
	})
	if err != nil {
		return nil
	}
	return controller
}
func (u *UserController) addUser(obj interface{}) {
	user := obj.(*apisUerV1.User)
	glog.Infof("User %s add", user.Name)
	_, exists, err := u.userFactory.Cnos().V1().Users().Informer().GetIndexer().GetByKey(user.Name)
	if err != nil {
		glog.Infof("get user error:%v", err)
		return
	}
	if exists {
		glog.Errorf("user %s is already exists,please choose another userName", user.Name)
		return
	}
	if strings.EqualFold(user.Name, "admin") || strings.EqualFold(user.Spec.Username, "admin") {
		glog.Errorf("user's name like admin and Admin is not allowed,please choose another name!")
		return
	}
	err = u.generateUserSecret(user)
	if err != nil {
		glog.Errorf("failed to create user %s secret ,error:%v", user.Name, err)
		return
	}
	err = u.createUser(user)
	if err != nil {
		glog.Errorf("create user failed,error:%v", err)
		return
	}
	glog.Infof("create user success!")
}
func (u *UserController) updateUser(oldObj, newObj interface{}) {
	oldUser, newUser := oldObj.(*apisUerV1.User), newObj.(*apisUerV1.User)
	if newUser.Spec.Username != oldUser.Spec.Username || newUser.Spec.AdminRole != oldUser.Spec.AdminRole ||
		newUser.Spec.Namespace != oldUser.Spec.Namespace {
		glog.Errorf("user's username or AdminRole or namespace can not modify!")
		return
	}
	fmt.Println("old password:", oldUser.Spec.Password)
	fmt.Println("new password", newUser.Spec.Password)
	//TODO 更新失败？
	if oldUser != newUser {
		newUss, errs := u.userClientSet.CnosV1().Users().Get(context.TODO(), newUser.Name, metav1.GetOptions{})
		if errs != nil {
			glog.Errorf("failed to get user", errs)
		}
		fmt.Println(newUss.Name, newUss.Spec.Username)
		_, err := u.userClientSet.CnosV1().Users().Update(context.TODO(), newUser, metav1.UpdateOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				fmt.Println("is not found")
				return
			}
			glog.Errorf("update user error:%v", err)
			return
		}
	}
	fmt.Println(oldUser.Spec.Username, "||||", newUser.Spec.Namespace)
	if newUser.Spec.AdminRole == true {
		if oldUser.Spec.Enabled == true && newUser.Spec.Enabled == false {
			err := u.deleteClusterRoleBinding(newUser)
			if err != nil {
				glog.Errorf("in updateUser delete clusterRoleBinding err %v", err)
				return
			}
			newUser.Status.Status = corev1.ConditionFalse
		}
		if oldUser.Spec.Enabled == false && newUser.Spec.Enabled == true {
			err := u.createClusterRoleBinding(newUser)
			if err != nil {
				glog.Errorf("in updateUser create clusterRoleBinding err:%v", err)
				return
			}
			newUser.Status.Status = corev1.ConditionTrue
		}
	} else { //normal user
		if oldUser.Spec.Enabled == true && newUser.Spec.Enabled == false {
			err := u.deleteRoleBinding(newUser)
			if err != nil {
				glog.Errorf("in updateUser delete rolebinding err %v", err)
				return
			}
			newUser.Status.Status = corev1.ConditionFalse
		}
		if oldUser.Spec.Enabled == false && newUser.Spec.Enabled == true {
			var role *rbacV1.Role
			var err error
			//if roles not change
			if util.CompareSlice(oldUser.Spec.Roles, newUser.Spec.Roles) {
				role, err = u.kubeClientSet.RbacV1().Roles(newUser.Spec.Namespace).Get(context.TODO(),
					newUser.Spec.Username, metav1.GetOptions{})
				if err != nil {
					glog.Errorf("in updateUser get role err:%v", err)
					return
				}
			} else {
				// if role change create new role
				role, err = u.updateRole(oldUser, newUser)
				if err != nil {
					glog.Errorf("in updateUser update role err:%v", err)
					return
				}
			}
			_, err = u.createRoleBinding(role, newUser)
			if err != nil {
				glog.Errorf("in updateUser create role err:%v", err)
				return
			}
			newUser.Status.Status = corev1.ConditionTrue
		}
		if !util.CompareSlice(oldUser.Spec.Roles, newUser.Spec.Roles) {
			_, err := u.updateRole(oldUser, newUser)
			if err != nil {
				glog.Errorf("in updateUser updateRole err:%v", err)
				return
			}
		}
	}

	//u.patchUserStatus(newUser)
}
func (u *UserController) Run(stopCh <-chan struct{}) error {
	glog.Infof("user controller is starting...")
	go u.userFactory.Start(stopCh)
	if ok := cache.WaitForCacheSync(stopCh, u.userSynced); !ok {
		return fmt.Errorf("failed to wati for caches to sync")
	}
	glog.Infof("user controller is synced and waiting for events...")
	<-stopCh
	glog.Infof("user controller is stopping...")
	return nil
}
func (u *UserController) generateUserSecret(user *apisUerV1.User) error {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(user.Spec.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("faile to generatepassword hash:%v", err)
	}
	// create the secret object
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      user.Spec.Username,
			Namespace: user.Spec.Namespace,
		},
		StringData: map[string]string{
			"userPassword": string(PasswordHash),
			"username":     user.Spec.Username,
		},
	}

	// create or update the secret
	_, err = u.kubeClientSet.CoreV1().Secrets(user.Spec.Namespace).Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			_, err = u.kubeClientSet.CoreV1().Secrets(user.Spec.Namespace).Update(context.Background(), secret, metav1.UpdateOptions{})
			fmt.Println("update secret success!")
		} else if err != nil {
			return fmt.Errorf("failed to create/update secret for user %s/%s: %v", user.Namespace, user.Name, err)
		}
	} else {
		fmt.Println("create secret success!")
	}
	Secret, err := u.kubeClientSet.CoreV1().Secrets(user.Spec.Namespace).Get(context.Background(), secret.Name, metav1.GetOptions{})
	user.Spec.SecretName = Secret.Name
	return nil
}

func (u *UserController) createUser(user *apisUerV1.User) error {
	switch {
	case user.Spec.Username == "":
		return errors.New("user.Spec.Username can not be nil")
	case user.Spec.Password == "":
		return errors.New("user.Spec.Password can not be nil")
	case user.Spec.Namespace == "":
		return errors.New("user.Spec.Namespace can not be nil")
	}
	password, err := util.GetPassword(context.TODO(), user.Spec.SecretName, u.kubeClientSet.CoreV1().Secrets(user.Spec.Namespace))
	if err != nil {
		glog.Errorf("get user.spec.password failed ,error:%v", err)
		return err
	}
	user.Status.Password = password
	user.Spec.Password = ""
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	user.Status.PasswordUpdateTime = updateTime
	//create certificate file
	if err := u.createCertificate(user); err != nil {
		glog.Errorf("create Certificate error:%v", err)
		return err
	}
	glog.Infof("create ca success")
	//TODO create k8sUser(√)
	if err = u.createK8sUser(user); err != nil {
		glog.Errorf("create k8sUser error:%v", err)
	}
	glog.Infof("create k8sUser success!")
	//TODO create namespace(√)
	_, err = u.kubeClientSet.CoreV1().Namespaces().Get(context.TODO(), user.Spec.Namespace, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: user.Spec.Namespace,
				},
			}
			_, err = u.kubeClientSet.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
			if err != nil && !apierrors.IsAlreadyExists(err) {
				glog.Errorf("can not ensure namespace %s already create", user.Spec.Namespace)
			}
		} else {
			glog.Errorf("unable to determine if namespace %s already created :%v", user.Spec.Namespace, err)
			return err
		}
	}
	glog.Infof("create ns success!")
	//TODO role and rolebinding
	if user.Spec.AdminRole == true {
		_, err := u.kubeClientSet.RbacV1().ClusterRoleBindings().Get(context.TODO(), ClusterRole, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				glog.Errorf("can not find clusterRole %s,error:%v", ClusterRole, err)
				return err
			} else {
				glog.Errorf("get cluster role error:%v", err)
				return err
			}
		} else {
			if user.Spec.Enabled == true {
				err := u.createClusterRoleBinding(user)
				if err != nil {
					return err
				}
				user.Status.Status = corev1.ConditionTrue
			} else {
				user.Status.Status = corev1.ConditionFalse
			}
		}
		user.Status.Type = AdminUser
	} else {
		role, err := u.createRole(user)
		if err != nil {
			return err
		}
		glog.Infof("create role %s success", role.Name)
		if user.Spec.Enabled == true {
			roleBinding, err := u.createRoleBinding(role, user)
			if err != nil {
				return err
			}
			glog.Infof("create rolebinding %s success", roleBinding.Name)
			user.Status.Status = corev1.ConditionTrue
		} else {
			user.Status.Status = corev1.ConditionFalse
		}
		user.Status.Type = NormalUser
	}
	//err = u.patchUserStatus(user)
	//if err != nil {
	//	return err
	//}
	return nil
}
func (u *UserController) createCertificate(user *apisUerV1.User) error {
	user.Status.AuthFile = make(map[string]string)
	dir := filepath.Join("/tmp", user.Spec.Username)
	if err := os.MkdirAll(dir, 0777); err != nil {
		glog.Errorf("create Certificate directory error:%v", err)
		return err
	}
	glog.Infof("create directory success!")
	//create key
	key := filepath.Join(dir, user.Spec.Username+".key")
	//openssl genrsa -out user.key 2048
	keyCmd := []string{"openssl", "genrsa", "-out", key, "2048"}
	keyCmdOut, keyCmdErr := exec.Command(keyCmd[0], keyCmd[1:]...).CombinedOutput()
	if keyCmdErr != nil {
		glog.Errorf("openssl genrsa key failed,error:%v,output:%s", keyCmdErr, keyCmdOut)
		return keyCmdErr
	}
	user.Status.AuthFile[UserKeyFile] = key
	glog.Infof("create user %s key success", user.Spec.Username)
	//create csr
	csr := filepath.Join(dir, user.Spec.Username+".csr")
	cn := fmt.Sprintf("/CN=%s", user.Spec.Username)
	csrCmd := []string{"openssl", "req", "-new", "-key", key, "-out", csr, "-subj", fmt.Sprintf("%s", cn)}
	// openssl req -new -key user.key -out user.user.csr  -subj  CN=user
	csrCmdOut, csrCmdErr := exec.Command(csrCmd[0], csrCmd[1:]...).CombinedOutput()
	if csrCmdErr != nil {
		glog.Errorf("openssl gen csr error:%v,output:%s", csrCmdErr, csrCmdOut)
		return csrCmdErr
	}
	user.Status.AuthFile[UserCsrFile] = csr
	glog.Infof("create user %s csr success", user.Spec.Username)
	//create crt
	if !util.CheckFileExists(CACRT) || !util.CheckFileExists(CAKEY) {
		return errors.New("k8s ca or key file is not exists")
	}
	// openssl x509 -req -in user.csr -CA k8s.crt -CAkey k8s.key -CAcreateserial -out user.crt -days
	crt := filepath.Join(dir, user.Spec.Username+".crt")
	crtCmd := []string{"openssl", "x509", "-req", "-in", csr, "-CA",
		CACRT, "-CAkey", CAKEY, "-CAcreateserial", "-out", crt, "-days", DAY}
	crtCmdOut, crtCmdErr := exec.Command(crtCmd[0], crtCmd[1:]...).CombinedOutput()
	if crtCmdErr != nil {
		glog.Errorf("openssl gen crt error:%v,output:%s", crtCmdErr, crtCmdOut)
		return crtCmdErr
	}
	user.Status.AuthFile[UserCrtFile] = crt
	glog.Infof("create user %s crt success", user.Spec.Username)

	//u.patchUserStatus(user)
	if !util.CheckFileExists(key) || !util.CheckFileExists(csr) || !util.CheckFileExists(crt) {
		return errors.New("key or csr or crt file not exists")
	}
	return nil
}
func (u *UserController) createK8sUser(user *apisUerV1.User) error {
	key := user.Status.AuthFile[UserKeyFile]
	crt := user.Status.AuthFile[UserCrtFile]
	if !util.CheckFileExists(key) {
		return fmt.Errorf(" user's  key %s file not found", key)
	}
	if !util.CheckFileExists(crt) {
		return errors.New("user's crt  file not exists")
	}
	crtPath := fmt.Sprintf("--client-certificate=%s", crt)
	keyPath := fmt.Sprintf("--client-key=%s", key)
	cmd := exec.Command("kubectl", "config", "set-credentials", user.Spec.Username, crtPath, keyPath, "--embed-certs=true")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create k8sUser : %v,output : %s", err, string(output))
	}
	configCmd := exec.Command("kubectl", "config", "get-clusters")
	configPut, err := configCmd.CombinedOutput()
	if err != nil {
		glog.Errorf("failed to get k8s name:", err)
	}
	clusterName := fmt.Sprintf(string(configPut[5:]))
	cluster := fmt.Sprintf("--cluster=%s", clusterName)
	users := fmt.Sprintf("--user=%s", user.Spec.Username)
	userCmd := exec.Command("kubectl", "config", "set-context", user.Spec.Username, cluster, users)
	_, err = userCmd.CombinedOutput()
	if err != nil {
		glog.Errorf("failed to set context for user ,error:%v", err)
	}
	glog.Infof("create k8sUser success,and output is %s", string(output))
	return nil
}
func (u *UserController) createRole(user *apisUerV1.User) (*rbacV1.Role, error) {
	verbs := user.Spec.Roles
	if len(user.Spec.Roles) == 0 {
		verbs = []string{"get", "list", "watch"}
	}
	userRole := &rbacV1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      user.Spec.Username,
			Namespace: user.Spec.Namespace,
			Labels: map[string]string{
				"group": "cnos.inspur.com",
			},
		},
		Rules: []rbacV1.PolicyRule{{
			APIGroups: []string{"*"},
			Resources: []string{"*"},
			Verbs:     verbs,
		}},
	}
	role, err := u.kubeClientSet.RbacV1().Roles(user.Spec.Namespace).Create(context.TODO(), userRole, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			glog.Infof("role %s,is exists", role.Name)
			return role, nil
		}
		glog.Errorf("failed to create role %v", err)
		return nil, err
	}
	glog.Infof("create role for user %s in namespace %s", user.Spec.Username, user.Spec.Namespace)
	return role, nil
}
func (u *UserController) updateRole(oldUser, newUser *apisUerV1.User) (*rbacV1.Role, error) {
	err := u.deleteRole(oldUser)
	if err != nil {
		return nil, err
	}
	role, err := u.createRole(newUser)
	if err != nil {
		return nil, err
	}
	glog.Infof("update role success")
	return role, nil
}
func (u *UserController) deleteRole(user *apisUerV1.User) error {
	if err := u.roleExists(user); err != nil {
		return err
	}
	err := u.kubeClientSet.RbacV1().Roles(user.Spec.Namespace).Delete(context.TODO(), user.Spec.Username, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	glog.Infof("delete role success")
	return nil
}
func (u *UserController) createRoleBinding(role *rbacV1.Role, user *apisUerV1.User) (*rbacV1.RoleBinding, error) {
	_, err := u.kubeClientSet.RbacV1().Roles(user.Spec.Namespace).Get(context.TODO(), role.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get role %s:%v", role.Name, err)
	}
	roleBinding := &rbacV1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      user.Spec.Username,
			Namespace: user.Spec.Namespace,
			Labels: map[string]string{
				"group": "cnos.inspur.com",
			},
		},
		Subjects: []rbacV1.Subject{{
			Kind:     "User",
			APIGroup: rbacV1.GroupName,
			Name:     user.Spec.Username,
		}},
		RoleRef: rbacV1.RoleRef{
			APIGroup: rbacV1.GroupName,
			Kind:     "Role",
			Name:     role.Name,
		},
	}
	userRoleBinding, err := u.kubeClientSet.RbacV1().RoleBindings(user.Spec.Namespace).Create(context.TODO(), roleBinding, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			glog.Infof("rolebinding %s is already exists", roleBinding.Name)
			return nil, err
		}
		glog.Errorf("failed to create RoleBinding %v", err)
		return nil, err
	}
	fmt.Println("create rolebinding :", roleBinding.Name, "success!")
	return userRoleBinding, nil
}
func (u *UserController) updateRoleBinding(oldUser, newUser *apisUerV1.User, oldRole, newRole *rbacV1.Role) error {
	roleBinding, err := u.kubeClientSet.RbacV1().RoleBindings(oldUser.Spec.Namespace).Get(context.TODO(), oldUser.Spec.Username, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("user %s exists but user's RoleBinding not exists:%v", oldUser.Spec.Username, err)
		}
		return fmt.Errorf("failed to get RoleBinding for user %s:%v", oldUser.Spec.Username, err)
	}
	if oldRole.Name != newRole.Name {
		err = u.deleteRoleBinding(oldUser)
		if err != nil {
			return fmt.Errorf("failed to delete RoleBinding for user %s:%v", oldUser.Spec.Username, err)
		}
		_, err = u.createRoleBinding(newRole, oldUser)
		if err != nil {
			return fmt.Errorf("failed to create RoleBinding for user %s:%v", newUser.Spec.Username, err)
		}
	} else {
		roleBinding.RoleRef.Name = newRole.Name
		_, err = u.kubeClientSet.RbacV1().RoleBindings(oldUser.Spec.Namespace).Update(context.Background(), roleBinding, metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("failed to update roleBinding for user %s:%v", oldUser.Spec.Username, err)
		}
	}
	glog.Infof("update rolebinding success")
	return nil
}
func (u *UserController) deleteRoleBinding(user *apisUerV1.User) error {
	roleBinding, err := u.kubeClientSet.RbacV1().RoleBindings(user.Spec.Namespace).Get(context.TODO(), user.Spec.Username, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return errors.New(fmt.Sprintf("user %s exists but user's RoleBinding not exists %v", user.Spec.Username, err))
		} else {
			return errors.New(fmt.Sprintf("in delete RoleBinding get roleBinding error:%v", err))
		}
	}

	err = u.kubeClientSet.RbacV1().RoleBindings(user.Spec.Namespace).Delete(context.TODO(), roleBinding.Name, metav1.DeleteOptions{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete RoleBinding %s:%v", roleBinding, err))
	}
	glog.Infof("delete rolebinding success")
	return nil
}
func (u *UserController) createClusterRoleBinding(user *apisUerV1.User) error {
	clusterRoleBinding := &rbacV1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: user.Spec.Username,
		},
		Subjects: []rbacV1.Subject{{
			Kind:     "User",
			APIGroup: rbacV1.GroupName,
			Name:     user.Spec.Username,
		}},
		RoleRef: rbacV1.RoleRef{
			APIGroup: rbacV1.GroupName,
			Kind:     "ClusterRole",
			Name:     ClusterRole,
		},
	}
	_, err := u.kubeClientSet.RbacV1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, metav1.CreateOptions{})
	if err != nil {
		glog.Errorf("failed to create clusterRoleBinding,error:%v", err)
		return err
	}
	glog.Infof("create cluster rolebinding success!")
	return nil
}
func (u *UserController) deleteClusterRoleBinding(user *apisUerV1.User) error {
	_, err := u.kubeClientSet.RbacV1().ClusterRoleBindings().Get(context.TODO(), user.Name, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return errors.New(fmt.Sprintf("user %s exists but user's clusterRoleBinding not exists %v", user.Spec.Username, err))
		} else {
			return errors.New(fmt.Sprintf("in deleteClusterRoleBinding get clusterrolebindings error %v", err))
		}
	}
	err = u.kubeClientSet.RbacV1().ClusterRoleBindings().Delete(context.TODO(), user.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (u *UserController) patchUserStatus(user *apisUerV1.User) error {
	bytess, err := user.Status.Bytes()
	if err != nil {
		glog.Error(err)
		return err
	}
	if _, err = u.userClientSet.CnosV1().Users().Patch(context.TODO(), user.Name,
		types.MergePatchType, bytess, metav1.PatchOptions{}, "status"); err != nil {
		glog.Errorf("patch user status failed", err)
		return err
	}
	return nil
}
func (u *UserController) roleExists(user *apisUerV1.User) error {
	_, err := u.kubeClientSet.RbacV1().Roles(user.Spec.Namespace).Get(context.TODO(), user.Spec.Username, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("user %s exists but user's role is not exists", user.Spec.Username)
		}
		return fmt.Errorf("error get role for user %s:%v", user.Spec.Username, err)
	}
	return nil
}
