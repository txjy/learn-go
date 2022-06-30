package routes

import (
	v1 "go_bin_blog/api/v1"
	"go_bin_blog/middleware"
	"go_bin_blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	{
		//用户模块的路由接口

		auth.PUT("user/:id", v1.UpdateUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)

		auth.PUT("category/:id", v1.UpdateCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.UpdateArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	router := r.Group("api/v1")
	{
		router.GET("users", v1.GetUsers)
		router.POST("user/add", v1.AddUser)
		router.GET("category", v1.GetCate)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)

	}

	r.Run(":3000")
	// r.Run(utils.HttpPort)
}
