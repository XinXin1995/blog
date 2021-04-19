package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBlogCommentRouter(router *gin.RouterGroup)  {
	BaseRouter := router.Group("comment")
	{
		BaseRouter.POST("list", v1.GetCommentList)
	}
}
