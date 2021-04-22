package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("category")
	{
		BaseRouter.GET("list", v1.GetCategories)
		BaseRouter.GET("statistic", v1.GetCategoryStatistic)
	}
}
