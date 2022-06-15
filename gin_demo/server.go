package main

import (
	"demo/controllers"
	"demo/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.Use(middleware.MyAuth())

	server.GET("/ping", func(context *gin.Context) {

		context.String(200, "%s", "pong")
	})

	server.Static("/resources", "./resources")
	server.StaticFile("/mydemoviedo", "./resources/demo.mp4")

	videocontroller := controllers.NewVideoController()
	videoGroup := server.Group("./videos")
	videoGroup.Use(middleware.MyLogger())

	//GET /videos
	videoGroup.GET("./", videocontroller.GetAll)

	//PUT /videos/1234
	videoGroup.PUT("./:id", videocontroller.Update)

	//POST /videos
	videoGroup.POST("./", videocontroller.Create)

	//DELETE /videos/123
	videoGroup.DELETE("./:id", videocontroller.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
