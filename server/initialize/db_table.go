package initialize

import (
	"blog/global"
	"blog/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.SysUser{},
		&model.SysMenu{},
		&model.BlogArticle{},
		&model.BlogTag{},
		&model.BlogCategory{},
		&model.SysFile{},
)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
