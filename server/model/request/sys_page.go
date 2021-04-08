package request

type SysAddPage struct {
	Name          string `json:"name"  binding:"required"`
	Path          string `json:"path"  binding:"required"`
	RedirectPath  string `json:"redirectPath"`
	Exec          bool   `json:"exec"`
	ComponentPath string `json:"componentPath" binding:"required"`
	Button        string `json:"button"`
	ParentId      int64  `json:"parentId"`
}

type SysUpdatePage struct {
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name"  binding:"required"`
	Path          string `json:"path"  binding:"required"`
	RedirectPath  string `json:"redirectPath"`
	Exec          bool   `json:"exec" `
	ComponentPath string `json:"componentPath" binding:"required"`
	Button        string `json:"button"`
	ParentId      int64  `json:"parentId"`
}

type SysDelPage struct {
	Ids []int64 `json:"ids" binding:"required"`
}