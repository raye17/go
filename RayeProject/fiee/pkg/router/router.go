package router

import (
	"test001/ssh/seller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors(), middleware.GinRecovery(true))
	//r.Use(middleware.CheckLogin(service.AccountProvider))

	r.POST("/report", seller.UpdateReportList)
	return r
}
