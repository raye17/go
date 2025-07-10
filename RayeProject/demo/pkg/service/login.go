package service

import (
	"fmt"
	"strconv"
	"sxy/demo/pkg/cache"
	"sxy/demo/pkg/common/jwt"
	"sxy/demo/pkg/db/model"
	"sxy/demo/pkg/internal/controller"
	"time"

	"sxy/demo/pkg/model/login"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req login.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, "参数错误")
		return
	}
	if req.Username == "" || req.Password == "" {
		Error(c, "用户名或密码不能为空")
		return
	}
	if err := controller.CreateUser(model.User{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		Error(c, "注册失败", err)
		return
	}
	Success(c, "注册成功", req.Username)
}
func Login(c *gin.Context) {
	ip := c.ClientIP()
	var req login.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, "参数错误")
		return
	}
	userinfo, err := controller.GetUserInfoByUsername(req.Username)
	if err != nil {
		Error(c, "获取用户信息失败", err)
		return
	}
	// 获取失败次数
	failKey := fmt.Sprintf(login.FailLoginKey, ip, req.Username)
	failCount := 0
	failCountStr, err := cache.RedisClient.Get(failKey).Result()
	if err == redis.Nil {
		failCount = 0
	} else if err != nil {
		Error(c, "Redis 获取失败次数出错", err)
		return
	} else {
		failCount, _ = strconv.Atoi(failCountStr)
	}

	if failCount >= 5 {
		Error(c, "登录失败次数过多,请5分钟后再试")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(req.Password)); err != nil {
		cache.RedisClient.Incr(failKey)
		if failCount == 0 {
			cache.RedisClient.Expire(failKey, 5*time.Minute)
		}
		Error(c, "密码错误", err)
		return
	}
	cache.RedisClient.Del(failKey)
	token, err := jwt.GenerateToken(int(userinfo.Id), userinfo.Username)
	if err != nil {
		Error(c, "生成token失败", err)
		return
	}
	Success(c, "登录成功", token)
}
func GetUserInfo(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, "参数错误")
		return
	}
	userInfo, err := controller.GetUserInfoByID(req.ID)
	if err != nil {
		Error(c, "获取信息失败", err)
		return
	}
	info := login.GetUserInfoFromC(c)
	data := login.UserInfo{
		UserId:   int(userInfo.Id),
		Username: userInfo.Username,
		Token:    info.Token,
	}
	Success(c, "获取信息成功", data)
}
func Logout(c *gin.Context) {
	userInfo := login.GetUserInfoFromC(c)
	claims, err := jwt.ParseToken(userInfo.Token)
	if err != nil {
		Error(c, "token解析失败", err)
		return
	}
	expiresAt := time.Until(time.Unix(claims.ExpiresAt, 0))
	key := fmt.Sprintf(login.BlackLoginKey, userInfo.UserId, userInfo.Username)
	err = cache.SetBlackLoginKey(key, userInfo.Token, expiresAt)
	if err != nil {
		Error(c, "缓存token失败", err)
		return
	}
	Success(c, "退出成功", nil)
}
