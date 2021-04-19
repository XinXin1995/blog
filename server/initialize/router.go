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
	{
		router.InitUserRouter(ApiGroup)
		router.InitCaptchaRouter(ApiGroup)
		router.InitArticleRouter(ApiGroup)
		router.InitTagRouter(ApiGroup)
		router.InitCategoryRouter(ApiGroup)
		router.InitManagerRouter(ApiGroup)
		router.InitFile(ApiGroup)
		router.InitBlogCommentRouter(ApiGroup)
	}
	//router.InitBase(ApiGroup)
	//router.InitPage(ApiGroup)
	return Router
}
