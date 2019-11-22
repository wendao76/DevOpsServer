package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserService struct {

}
func NewUser() *UserService{
	return & UserService {}
}

func (us * UserService) Get(c *gin.Context) {
	fmt.Println("GET")
}

func (us * UserService) Add(c *gin.Context) {
	fmt.Println("GET")
}

func (us * UserService) Remove(c *gin.Context) {
	fmt.Println("GET")
}

func (us * UserService) List(c *gin.Context) {
	fmt.Println("GET")
}
