package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("base")
	{
		UserRouter.GET("captcha", v1.Captcha)
	}
}
