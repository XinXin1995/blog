package model

import "time"

type SysRole struct {
	Id        string    `json:"id" xorm:"pk"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}
