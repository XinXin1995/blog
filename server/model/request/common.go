package request

type PageInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}
