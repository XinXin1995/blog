package model

import "blog/global"

type BlogArticle struct {
	global.GVA_MODEL
	Title      string       `json:"title"`
	CategoryId uint         `json:"categoryId"`
	Category   BlogCategory `json:"category"`
	Content    string       `json:"content" gorm:"type:text"`
	Thumb      string       `json:"thumb"`
	View       int64        `json:"view"`
	Likes      int64        `json:"likes"`
	TagIdGroup   string 	`json:"-"`
	Tags       []BlogTag    `json:"tags" gorm:"many2many:article_tags"`
}

type Articles struct {
	global.GVA_MODEL
	Title        string  `json:"title"`
	CategoryId   uint    `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	Content      string  `json:"content" gorm:"type:text"`
	Thumb        string  `json:"thumb"`
	View         int64   `json:"view"`
	Likes        int64   `json:"likes"`
	GroupTagId   string  `json:"groupTagId"`
	Tags         []BlogTag `json:"tags"`
}
