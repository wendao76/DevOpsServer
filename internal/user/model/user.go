package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string
    Avatar string
    Phone string
    Email string
    Verified int8
}
