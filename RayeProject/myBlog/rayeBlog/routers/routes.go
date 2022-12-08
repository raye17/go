package routers

import (
	"github.com/gin-gonic/gin"
	"rayeBlog/controller"
	"rayeBlog/middleware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	//允许跨域访问
	r.Use(middleware.CORSMiddleware())
	//注册
	r.POST("/register", controller.Register)
	//登录
	r.POST("/login", controller.Login)
	//上传图像
	r.POST("/upload", controller.Upload)

	//查询分类
	r.GET("/category", controller.SearchCategory)         //查询分类
	r.GET("/category/:id", controller.SearchCategoryName) //查询分类名

	//用户信息管理
	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthMiddleWare())
	userRoutes.GET("", controller.GetInfo)                     //获取用户信息
	userRoutes.GET("briefInfo/:id", controller.GetBriefInfo)   //获取用户简要信息
	userRoutes.GET("detailInfo/:id", controller.GetDetailInfo) //获取用户详细信息
	userRoutes.PUT("avatar/:id", controller.ModifyAvatar)      //修改头像
	userRoutes.PUT("name/:id", controller.ModifyName)          //修改用户名

	// 我的收藏
	colRoutes := r.Group("/collects")
	colRoutes.Use(middleware.AuthMiddleWare())
	colRoutes.GET(":id", controller.Collects)        // 查询收藏
	colRoutes.PUT("new/:id", controller.NewCollect)  // 收藏
	colRoutes.DELETE(":index", controller.UnCollect) // 取消收藏
	// 我的关注
	folRoutes := r.Group("/following")
	folRoutes.Use(middleware.AuthMiddleWare())
	folRoutes.GET(":id", controller.Following)      // 查询关注
	folRoutes.PUT("new/:id", controller.NewFollow)  // 关注
	folRoutes.DELETE(":index", controller.UnFollow) // 取消关注

	//文章的增删改查
	articleRoutes := r.Group("/article")
	articleController := controller.NewArticleController()
	articleRoutes.POST("", middleware.AuthMiddleWare(), articleController.Create)
	articleRoutes.PUT(":id", middleware.AuthMiddleWare(), articleController.Update)
	articleRoutes.DELETE(":id", middleware.AuthMiddleWare(), articleController.Delete)
	articleRoutes.GET(":id", middleware.AuthMiddleWare(), articleController.Show)
	articleRoutes.POST("list", middleware.AuthMiddleWare(), articleController.List)
	return r
}
