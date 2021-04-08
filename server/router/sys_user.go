package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("add", v1.Register)
		UserRouter.GET("detail", v1.UserDetail)
		UserRouter.GET("list", v1.UserList)
		UserRouter.POST("valid", v1.Valid)
		UserRouter.POST("bindRole", v1.UserBindRole)
	}
}
