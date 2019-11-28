package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/internal/common/dao"
	"go_web/internal/common/util"
	"go_web/internal/user/model"
	"net/http"
)

type UserAction struct {
}

//用户表单结构体
type UserForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Email string	`form:"email" json:"email"`
	Phone string	`form:"phone" json:"phone" binding:"required"`
}

func NewUser() *UserAction{
	return & UserAction {}
}

//@Summary 获取用户信息
//@Description 根据用户ID获取用户信息
//@Accept json
//@Produce json
//@Router /users [get]
func (ua * UserAction) Get(ctx *gin.Context) {
	userID := ctx.Param("id")
	var userModel model.User
	db := dao.Get().Db

	db.First(&userModel, userID)
	if userModel.ID == 0 {
		errors(ctx, http.StatusBadRequest, "指定用户不存在")
		return
	}
	result(ctx, userModel, http.StatusOK)
}

//@Summary 新增用户
//@Description 新增一个用户
//@Accept json
//@Produce json
//@Router /users [post]
func (ua * UserAction) Create(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.Bind(&userForm)
	if err != nil {
	    errors(ctx, http.StatusBadRequest, err.Error())
	    return
	}
	var user model.User
	util.CopyStruct(&user, &userForm)
	daoIns := dao.Get()
	db := daoIns.Db.Create(&user)
	if db.Error != nil {
	    errors(ctx, http.StatusServiceUnavailable, db.Error.Error())
	    return
	}
	result(ctx, nil, OK)
}

//@Summary 修改用户
//@Description 修改用户信息
//@Accept json
//@Produce json
//@Router /users [put]
func (ua *UserAction) Modify(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.Bind(&userForm)
	if err != nil {
		errors(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var user model.User

	db := dao.Get().Db
	db.First(&user)
	util.CopyStruct(&user, &userForm)
	db = db.Save(user)
	if db.Error != nil {
		errors(ctx, http.StatusServiceUnavailable, db.Error.Error())
		return
	}
	result(ctx, nil, OK)
}

//@Summary 删除用户
//@Description 根据用户ID删除用户
//@Accept json
//@Produce json
//@Router /users [delete]
func (ua * UserAction) Remove(ctx *gin.Context) {
	userID := ctx.Param("id")
	daoIns := dao.Get()
	var userObj model.User
	daoIns.Db.First(&userObj, userID)
	daoIns.Db.Delete(&userObj)
	result(ctx, nil, http.StatusOK)
}

func (ua * UserAction) List(c *gin.Context) {
	fmt.Println("List")
}
