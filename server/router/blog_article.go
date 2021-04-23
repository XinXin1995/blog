package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("article")
	{
		BaseRouter.POST("list", v1.GetArticleList)
		BaseRouter.GET("like", v1.LikeArticle)
		BaseRouter.GET("detail", v1.GetArticle)
		BaseRouter.GET("top", v1.SetTopArticle)
		/*BaseRouter.GET("union", controller.GetArticleUnion)
		BaseRouter.GET("like", controller.LikeArticle)*/
	}
}
