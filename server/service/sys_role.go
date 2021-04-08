package service

import (
	"blog/global"
	"blog/model"
)

func AddRole(r model.SysRole) (error, interface{}) {
	affected, err := global.GVA_DB.Insert(&r)
	return err, affected
}
