package request

type AddTag struct {
	Name string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}