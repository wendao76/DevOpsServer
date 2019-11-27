package http

import (
	"github.com/gin-gonic/gin"
	"go_web/internal/common/middleware"
	"log"
)

func New() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	initRouter(engine)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("用户服务启动失败!")
	}
}

//路由初始化
func initRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userServer := NewUser()
	userGroup := engine.Group("/users")
	{
		userGroup.POST("", middleware.Auth(),  userServer.Add)
		userGroup.GET("", userServer.List)
		userGroup.GET("/:userId", userServer.Get)
		userGroup.DELETE("/:userId", userServer.Remove)
	}
}
