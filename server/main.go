package main

import (
	"blog/core"
	"blog/global"
	"blog/initialize"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()

	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()
	core.RunWindowsServe()
}
