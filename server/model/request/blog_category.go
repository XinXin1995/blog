package request

type AddCategory struct {
	Name string `json:"name" binding:"required"`
}