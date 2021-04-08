package global

import (
	"blog/config"
	"github.com/go-xorm/xorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *xorm.Engine
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *oplogging.Logger
)
