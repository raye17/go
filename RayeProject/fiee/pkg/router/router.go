package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors(), middleware.GinRecovery(true))
	//r.Use(middleware.CheckLogin(service.AccountProvider))
	return r
}
