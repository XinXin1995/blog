package request

type PageInfo struct {
	PageNo     int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}
