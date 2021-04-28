package request

type AddArticleParam struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category uint   `json:"category" binding:"required"`
	Tags     []uint `json:"tags" binding:"required"`
	Thumb    string `json:"thumb"`
}

type UpdateArticleParam struct {
	ID       uint   `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category uint   `json:"category" binding:"required"`
	Tags     []uint `json:"tags" binding:"required"`
	Thumb    string `json:"thumb" binding:"required"`
}

type ArticleList struct {
	PageInfo
	Tags        []int `json:"tags"`
	Category    int   `json:"category"`
	OrderType   int   `json:"orderType"`
	WithContent bool  `json:"withContent"`
}
