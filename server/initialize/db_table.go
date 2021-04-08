package initialize

import (
	"blog/global"
	"blog/model"
)

func DBTables() {
	db := global.GVA_DB
	db.Sync2(
		&model.SysUser{},
		&model.SysUserRole{},
		&model.SysRole{},
		&model.SysPage{},
	)
}
