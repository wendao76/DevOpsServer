package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/internal/common/dao"
	"go_web/internal/user/model"
	"time"
)

type UserController struct {

}
func NewUser() *UserController{
	return & UserController {}
}

func (us * UserController) Get(ctx *gin.Context) {
	fmt.Println("GET")
}

//@Summary 新增用户
//@Description 新增一个用户
//@Accept json
//@Produce json
//@Router /users [get]
func (us * UserController) Add(ctx *gin.Context) {
	user := &model.User {
	    Username: "test",
	    Password : "test",
	    Nickname: "wendao",
	    Email: "tiger900721@163.com",
	    Phone: "13255055666",
	}

	daoIns := dao.Get()

	daoIns.Db.Create(user)
	redisClient := daoIns.Redis

	err := redisClient.Set("test-redis-key", "fasdfasdf", 300 * time.Second).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Add")
}

func (us * UserController) Remove(ctx *gin.Context) {
	fmt.Println("Remove")
}

func (us * UserController) List(c *gin.Context) {
	fmt.Println("List")
}
