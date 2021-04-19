package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("login", v1.Login)
		UserRouter.POST("visit", v1.LoginForVisitor)
	}
}
