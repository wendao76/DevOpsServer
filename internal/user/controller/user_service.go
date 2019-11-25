package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/internal/common/dao"
	"go_web/internal/user/model"
	"log"
	"time"
)

type UserService struct {

}
func NewUser() *UserService{
	return & UserService {}
}

func (us * UserService) Get(c *gin.Context) {
	fmt.Println("GET")
}

//@Summary 新增用户
//@Description 新增一个用户
//@Accept json
//@Produce json
//@Router /users
func (us * UserService) Add(c *gin.Context) {
	user := &model.User {
	    Username: "test",
	    Password : "test",
	    Nickname: "wendao",
	    Email: "tiger900721@163.com",
	    Phone: "13255055666",
	}

	dao, err := dao.Get()
	if err != nil {
		log.Fatal("dao获取失败")
	}
	db := dao.Db
	db.Create(user)

	redisClient := dao.Redis

	err = redisClient.Set("test-key", "fasdfasdf", 300 * time.Second).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Add")
}

func (us * UserService) Remove(c *gin.Context) {
	fmt.Println("Remove")
}

func (us * UserService) List(c *gin.Context) {
	fmt.Println("List")
}
