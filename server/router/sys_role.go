package router

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRole(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role").Use(middleware.JWTAuth())
	{
		RoleRouter.POST("add", v1.AddRole)
	}
}
