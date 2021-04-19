package request

type AddComment struct {
	ParentId uint   `json:"parentId"`
	ArticleId uint `json:"articleId" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
type TreeComment struct {
	PageInfo
	ArticleId uint `json:"articleId" binding:"required"`
}
