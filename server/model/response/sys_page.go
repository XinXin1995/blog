package response

type PageTree struct {
	Id            int64       `json:"id"`
	Name          string      `json:"name"`
	Path          string      `json:"path"`
	RedirectPath  string      `json:"redirectPath"`
	Exec          bool        `json:"exec"`
	ComponentPath string      `json:"componentPath"`
	ParentId      int64       `json:"parentId"`
	Children      []*PageTree `json:"children"`
}
