package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBase(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.GET("captcha/:captchaId", v1.CaptchaImg)
	}
}
