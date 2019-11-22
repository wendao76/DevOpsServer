package main

import (
	"github.com/gin-gonic/gin"
	"go_web/internal/common/config"
	"go_web/internal/common/dao"
	"go_web/internal/user/service"
	"log"
)

//@title 接口文档
//@version 0.0.1
func main() {
	StartServer()
}

//用户服务
func StartServer(){
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	initRouter(engine)
	err := config.InitDefault()
	conf, err := config.Default()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = dao.Init(conf)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine.Run(":8080")
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
	userService := service.NewUser()
	userGroup := engine.Group("/users")
	{
		userGroup.POST("", userService.Add)
		userGroup.GET("", userService.List)
		userGroup.GET("/:userId", userService.Get)
		userGroup.DELETE("/:userId", userService.Remove)
	}
}
