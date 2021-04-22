package response

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`
	Count int `json:"count"`
}
