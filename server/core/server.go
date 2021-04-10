package core

import (
	"blog/global"
	"blog/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServe() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
