package model

import "time"

// JS Router Config
type SysPage struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Path          string    `json:"path"`
	RedirectPath  string    `json:"redirectPath"`
	Exec          bool      `json:"exec"`
	ComponentPath string    `json:"componentPath"`
	Button        string    `json:"button"`
	ParentId      int64     `json:"parentId"`
	CreatedAt     time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt     time.Time `json:"createdAt" xorm:"updated"`
}
