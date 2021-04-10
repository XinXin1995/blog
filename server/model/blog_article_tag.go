package model

import (
	"blog/global"
)

type BlogTag struct {
	global.GVA_MODEL
	Name  string `json:"name"`
	Color string `json:"color"`
}
