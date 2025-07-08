package login

import "github.com/gin-gonic/gin"

const (
	BlackLoginKey = "blackLogin:%d-%s"
	LoginKey      = "login:%s"
	FailLoginKey  = "failLogin:%s:%s"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResp struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
}
type UserInfo struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func GetUserInfoFromC(c *gin.Context) UserInfo {
	userInfoAny, _ := c.Get("jwtInfo")
	userInfo := userInfoAny.(UserInfo)
	return userInfo
}
