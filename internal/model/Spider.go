package model

import "github.com/jinzhu/gorm"

//网站信息
type WebSite struct {
	gorm.Model
	Url string
	Name string `gorm:`
	Deep int64
	Md5 string
}

//抓取规则
type GrabRule struct {
	gorm.Model
	Name string
	WebSiteId int64
	Exp string
}