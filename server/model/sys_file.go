package model

import "blog/global"

type SysFile struct {
	global.GVA_MODEL
	Name   string `json:"name" gorm:"comment:文件名"`
	Url    string `json:"url" gorm:"comment:文件地址"`
	Prefix string `json:"prefix" gorm:"comment:文件存储前缀"`
	Tag    string `json:"tag" gorm:"comment:文件标签"`
	Key    string `json:"key" gorm:"comment:编号"`
}
