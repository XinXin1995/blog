package response

type StatisticCategory struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
