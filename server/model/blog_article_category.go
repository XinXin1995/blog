package model

import (
	"blog/global"
)

type BlogCategory struct {
	global.GVA_MODEL
	Name string `json:"name"`
	Type int    `json:"type" gorm:"comment: 分类类型"` // 1.个人简历  2.笔记  3.自己写的文章
}
