package service

import (
	"errors"
	"fmt"
	"raye/demo/db"
	"raye/demo/db/model"
	"raye/demo/pkg/cache"
	"raye/demo/pkg/utlis/e"
	"raye/demo/pkg/utlis/jwt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	loginCode = "sxy1017"
)

type login struct {
	Username string
	Password string
	Code     string
}
type UserInfo struct {
	Id uint
	login
	Age    int
	Gender string
	Token  string
}

func Login(c *gin.Context) {
	var req UserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
	}
	// 验证用户名和密码
	if req.Username == "" || req.Password == "" {
		fmt.Println("username or password is empty")
		ResponseMsg(c, e.Failed, "username or password is empty", nil, nil)
	}
	// 验证用户名和密码是否正确
	var user *model.User
	if err := db.DbTest01.Table("user").Where("name = ?", req.Username).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			fmt.Println("register user")
			u, errs := register(req)
			if errs != nil {
				ResponseMsg(c, e.Failed, errs.Error(), errs, nil)
				return
			}
			user = u
		} else {
			fmt.Println("get password error: ", err)
			ResponseMsg(c, e.Failed, err.Error(), err, nil)
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println("password error")
		ResponseMsg(c, e.Failed, e.PasswordError, errors.New("password error"), nil)
		return
	}
	// 生成JWT token
	token, err := jwt.GenerateToken(user.ID, user.Name, int(7*24), e.JWTSecret)
	if err != nil {
		fmt.Println("generate token error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	// 将token存储到Redis
	if err := cache.RedisClient.Set(c, fmt.Sprintf("user_token_%d", user.ID), token, time.Hour*24*7).Err(); err != nil {
		fmt.Println("set token to redis error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	fmt.Println("token:", token)
	ResponseMsg(c, e.Success, e.LoginSuccess, nil, token)
}
func register(req UserInfo) (user *model.User, err error) {
	// 验证用户名、密码和code
	if req.Username == "" || req.Password == "" {
		fmt.Println("username, password  is empty")
		return nil, errors.New("username, password is empty")
	}
	// 验证code是否正确
	// if req.Code != loginCode {
	// 	fmt.Println("code error")
	// 	return nil, errors.New("code error")
	// }
	// 验证用户名是否已经存在
	if err = db.DbTest01.Table("user").Where("name =?", req.Username).First(&user).Error; err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			fmt.Println("get password error: ", err)
			return
		}
	}
	if user != nil && user.ID != 0 {
		fmt.Println("username already exists")
		return
	}
	// 注册用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hash password error: ", err)
		return nil, err
	}
	var createReq = model.User{
		Name:     req.Username,
		Password: string(hashedPassword),
		Age:      req.Age,
		Gender:   req.Gender,
	}
	if err = db.DbTest01.Table("user").Create(&createReq).Error; err != nil {
		fmt.Println("create user error: ", err)
		return
	}
	return user, nil
}
func GetUserInfo(c *gin.Context) {
	var req UserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind json error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	var user model.User
	if err := db.DbTest01.First(&user, req).Error; err != nil {
		fmt.Println("get user by id error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, user)
}
func Logout(c *gin.Context) {
	// 从header中获取token
	token := c.GetHeader("Authorization")
	if token == "" {
		fmt.Println("get token error: token is empty")
		ResponseMsg(c, e.Failed, "token is empty", nil, nil)
		return
	}
	claims, err := jwt.ParseToken(token, e.JWTSecret)
	if err != nil {
		fmt.Println("parse token error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	// 将token加入Redis黑名单
	if err := cache.RedisClient.SAdd(c, "jwt:blacklist", token).Err(); err != nil {
		fmt.Println("add token to blacklist error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	// 设置黑名单token过期时间(基于token有效期)
	remainingTime := time.Until(time.Unix(claims.ExpiresAt, 0))
	if err := cache.RedisClient.Expire(c, "jwt:blacklist", remainingTime).Err(); err != nil {
		fmt.Println("set blacklist expire error: ", err)
		ResponseMsg(c, e.Failed, err.Error(), err, nil)
		return
	}
	ResponseMsg(c, e.Success, e.SuccessMsg, nil, "logout success")
}
func GetUserInfoDetail(c *gin.Context) (userInfo interface{}, err error) {
	// 从上下文中获取用户claims
	userInfo, exists := c.Get("userInfo")
	if !exists {
		fmt.Println("info:", userInfo)
		return nil, errors.New("user not found")
	}
	return userInfo, nil
}
