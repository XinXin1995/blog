package model

import "blog/global"

type SysMenu struct {
	global.GVA_MODEL
	Id       int64  `json:"id" xorm:"autoincr pk"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
	Router   string `json:"router"`
}
