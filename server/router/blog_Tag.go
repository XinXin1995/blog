package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("tag")
	{
		BaseRouter.GET("list", v1.GetTags)
	}
}
