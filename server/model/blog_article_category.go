package model

import (
	"blog/global"
)

type BlogCategory struct {
	global.GVA_MODEL
	Name string `json:"name"`
}
