package model

import "github.com/jinzhu/gorm"

//数据库对象


//用户表
type User struct {
	gorm.Model
	Username string
	Password string
	UserType uint
	Avatar string
	Email string
	Intro string
}

//角色表
type Role struct {
	gorm.Model
	Name string
	Code string
	Remarks string
}

//资源表
type Resource struct {
	gorm.Model
	Name string
	Code string
	PermitCode string
	Module string
}

//用户角色表
type UserRole struct {
	gorm.Model
	UserId uint
	RoleId uint
}

//角色资源表
type RoleResource struct {
	gorm.Model
	RoleId uint
	ResourceId uint
}


