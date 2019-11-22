package main

import (
	"github.com/gin-gonic/gin"
	"go_web/internal/common/config"
	"go_web/internal/user/service"
	"log"
)

func main() {
	StartServer()
}

//用户服务
func StartServer(){
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	initRouter(engine)
	err := engine.Run(":8081")
	if err != nil {
		log.Fatal("用户服务启动失败!")
	}
	err = config.InitDefault()
	if err != nil {
		log.Fatal("配置初始化失败")
	}
}

func initRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userService := service.NewUser()
	userGroup := engine.Group("/users")
	{
		userGroup.POST("", userService.Add)
		userGroup.GET("", userService.List)
		userGroup.GET("/:userId", userService.Get)
		userGroup.DELETE("/:userId", userService.Remove)
	}
}
