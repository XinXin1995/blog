package model

import (
	"blog/global"
)

type SysUser struct {
	global.GVA_MODEL
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     int    `json:"role" gorm:"default:3"` // 1：管理员 3：普通人员
}
