package model

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	Name string //名称
	Url  string //地址
	Sort int    `gorm:"default:'0'"` //排序
	View int    //访问次数
}
