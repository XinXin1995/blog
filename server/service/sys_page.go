package service

import (
	"blog/global"
	"blog/model"
	"fmt"
)

func AddPage(page *model.SysPage) (int64, error) {
	return global.GVA_DB.Insert(page)
}

func UpdatePage(page *model.SysPage) (int64, error) {
	return global.GVA_DB.Id(page.Id).Update(page)
}

func DeletePage(ids []int64) (int64, error) {
	var idStr string
	for i, id := range ids {
		if i == len(ids)-1 {
			idStr += fmt.Sprintf("%d", id)
		} else {
			idStr += fmt.Sprintf("%d,", id)
		}
	}
	return global.GVA_DB.Where("id in (?)", idStr).Delete(&model.SysPage{})
}

func AllPages() (pages []model.SysPage, err error) {
	err = global.GVA_DB.Find(&pages)
	return
}
