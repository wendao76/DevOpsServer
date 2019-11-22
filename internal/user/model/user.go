package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string
    Password string
    Nickname string
    Avatar string
    Phone string
    Email string
}
