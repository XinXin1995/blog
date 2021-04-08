package core

import (
	"blog/global"
	"blog/initialize"
	"fmt"
	"net/http"
	"time"
)

func RunWindowsServe() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	global.GVA_LOG.Error(s.ListenAndServe())
}
