package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("authentication"); token != "" {
			if passed := AuthByToken(token, ctx); !passed {
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}

//根据用户密码登录
func AuthByUsername(username string, password string, ctx *gin.Context) bool {
	return true
}

//根据token进行登录
func AuthByToken(token string, ctx *gin.Context) bool {
	fmt.Println("AuthByToken: token: " + token)
	return true
}

//根据小程序信息进行登录
func AuthByMina(loginData interface{}) bool{
	return true
}
