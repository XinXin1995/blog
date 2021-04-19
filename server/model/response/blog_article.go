package response

import (
	"blog/global"
	"blog/model"
)

type RespArticle struct {
	global.GVA_MODEL
	Title      string             `json:"title"`
	CategoryId uint               `json:"categoryId"`
	Category   model.BlogCategory `json:"category"`
	Thumb      string             `json:"thumb"`
	View       int64              `json:"view"`
	Likes      int64              `json:"likes"`
	Tags       []model.BlogTag    `json:"tags"`
}
