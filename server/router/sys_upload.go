package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFile(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("file")
	{
		BaseRouter.POST("upload", v1.Upload)
	}
}
