package router

import (
	"raye/demo/pkg/service"

	"github.com/gin-gonic/gin"
)

func ImgRouter(r *gin.RouterGroup) {
	auth := r.Group("")
	//auth.Use(middleware.JWTMiddleware())
	auth.POST("img/upload", service.UploadImg)
}
