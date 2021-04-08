package model

import (
	"time"
)

type SysUser struct {
	Id        string    `json:"id" xorm:"pk"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

type SysUserRole struct {
	Id        int64  `json:"id"`
	UserId    string `json:"userId"`
	RoleId    string `json:"roleId"`
	IsDefault bool   `json:"isDefault"`
}
