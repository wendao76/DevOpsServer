package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//客户端注册信息
type ClientModel struct {
	gorm.Model
	Name string
	Secret string
	Domain string
	UserID string
}
// GetID client id
func (c *ClientModel) GetID() string {
	return strconv.Itoa(int(c.ID))
}

// GetSecret client domain
func (c *ClientModel) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *ClientModel) GetDomain() string {
	return c.Domain
}

// GetUserID user id
func (c *ClientModel) GetUserID() string {
	return c.UserID
}





