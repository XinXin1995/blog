package response

import "time"

type RespArticle struct {
	ID        uint                    `json:"id"`
	Title     string                  `json:"title"`
	Category  RespArticleCategory `json:"category"`
	Content   string                  `json:"content"`
	Thumb     string                  `json:"thumb"`
	View      int64                   `json:"view"`
	Likes     int64                   `json:"likes"`
	Tags      []RespArticleTag    `json:"tags"`
	CreatedAt time.Time               `json:"createdAt"`
	UpdatedAt time.Time               `json:"updatedAt"`
}

type RespArticleCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type RespArticleTag struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
