package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/internal/common/exception"
)

type RespBase struct {
    Msg string	`json:"msg"`
    Code int	`json:"code"`
    Data interface {}	`json:"data,omitempty"`
}

type RespPage struct {
	*RespBase
}

type RespError struct {
    *RespBase
    ErrorCode string `json:"error_code"`
}

//返回空数据
func responseOk(ctx *gin.Context) {
	respBody := & RespBase{
		Code: 200,
		Msg: "ok",
	}
	ctx.JSON(200, respBody)
}

//返回分页数据
func responsePage(ctx *gin.Context) {

}

//TODO 异常返回
func responseError(code int, err exception.Exception, ctx *gin.Context) {
	ctx.Abort()
	e := & RespError{
		ErrorCode : err.Coder(),
	}

	fmt.Println(e)
	e.Code = code
	e.Msg = err.Error()
	ctx.JSON(500, e)
}
