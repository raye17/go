package router

import (
	"raye/demo/pkg/middleware"
	"raye/demo/pkg/service"

	"github.com/gin-gonic/gin"
)

func ImgRouter(r *gin.RouterGroup) {
	auth := r.Group("")
	noAuth := r.Group("")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.POST("img/upload", service.UploadImg)
		auth.POST("img/list", service.GetUploadedFiles)
	}
	{
		noAuth.POST("img/oss/upload", service.PutObject)
		noAuth.POST("img/oss/uploads", service.PutObjectByBytes)
		noAuth.POST("img/oss/list", service.ListObjects)
	}

}
