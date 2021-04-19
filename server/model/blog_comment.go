package model

import "blog/global"

type BlogComment struct {
	global.GVA_MODEL
	ArticleId  uint          `json:"articleId" gorm:"comment:文章ID"`
	ParentId   uint          `json:"parentId" gorm:"comment:回复评论ID"`
	Content    string        `json:"content" gorm:"comment:回复内容"`
	Username   string        `json:"username" gorm:"回复人员名称"`
	UserAvatar string        `json:"avatar" gorm:"comment:回复人头像"`
	Likes      int           `json:"likes" gorm:"点赞次数"`
	Children   []BlogComment `json:"children" gorm:"-"`
}
