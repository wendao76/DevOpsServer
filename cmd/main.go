package main

import (
	"github.com/gin-gonic/gin"
	"go_web/internal/user/http"
	"log"
)

//@title 接口文档
//@version 0.0.1
//@tag.name GO语言搭建的web脚手架
//@tag.description 快速搭建环境
func main() {
	StartServer()
}

//用户服务
func StartServer(){
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	initRouter(engine)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("用户服务启动失败!")
	}
}

func initRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userService := http.NewUser()
	userGroup := engine.Group("/users")
	{
		userGroup.POST("", userService.Add)
		userGroup.GET("", userService.List)
		userGroup.GET("/:userId", userService.Get)
		userGroup.DELETE("/:userId", userService.Remove)
	}
}
