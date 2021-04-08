package initialize

import (
	"blog/middleware"
	"blog/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("api")
	router.InitUserRouter(ApiGroup)
	router.InitRole(ApiGroup)
	router.InitBase(ApiGroup)
	router.InitPage(ApiGroup)
	return Router
}
