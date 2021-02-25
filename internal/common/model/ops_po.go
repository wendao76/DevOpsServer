package model

import "github.com/jinzhu/gorm"

//运维对象
//环境
type OpEnv struct {
	gorm.Model
	Name string
	Code string
	Remarks string
}

//版本信息
type OpVersion struct {
	gorm.Model
	No uint
	Version string
	Remarks string
}

//项目
type OpPrj struct {
	gorm.Model
	Name string
	Code string
	Remarks string
	GitUrl string
	Owner string
	OwnerId uint
	EnvId uint
	EnvCode string
}

type OpPrjVersion struct {
	gorm.Model
	Version string
	PrjCode string
}

//脚本
type OpScript struct {
	gorm.Model
	Name string
	Type uint
	Status uint
	Content string
	Version string
	EnvId uint
	EnvCode string
}

//部署文档
type OpDocument struct {
	gorm.Model
	Title string
	Content string
	Status uint
	EnvId uint
	EnvCode string
}
