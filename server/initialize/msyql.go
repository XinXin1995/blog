package initialize

import (
	"blog/global"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	admin := global.GVA_CONFIG.Mysql
	db, err := xorm.NewEngine("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config)
	if err != nil {
		global.GVA_LOG.Error("数据库启动异常", err)
	} else {
		global.GVA_DB = db
		global.GVA_DB.SetMaxIdleConns(admin.MaxIdleConns)
		global.GVA_DB.SetMaxOpenConns(admin.MaxOpenConns)
	}
}
