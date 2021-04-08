package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPage(Router *gin.RouterGroup) {
	PageRouter := Router.Group("page")
	{
		PageRouter.POST("add", v1.AddPage)
		PageRouter.POST("update", v1.UpdatePage)
		PageRouter.POST("del", v1.DeletePage)
		PageRouter.GET("tree", v1.PageTree)
	}
}
