package router

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitManagerRouter(router *gin.RouterGroup) {
	managerGroup := router.Group("manager")
	managerGroup.Use(middleware.JWTAuth())
	{
		userGroup := managerGroup.Group("user")
		{
			userGroup.GET("detail", v1.UserDetail)
			userGroup.GET("list", v1.UserList)
		}
		articleRouter := managerGroup.Group("article")
		{
			articleRouter.POST("add", v1.AddArticle)
			articleRouter.PUT("update", v1.UpdateArticle)
			articleRouter.DELETE("delete", v1.DelArticle)
		}
		tagGroup := managerGroup.Group("tag")
		{
			tagGroup.POST("create", v1.CreateTag)
			tagGroup.PUT("update", v1.UpdateTag)
			tagGroup.DELETE("delete", v1.DeleteTag)
		}
		categoryGroup := managerGroup.Group("category")
		{
			categoryGroup.POST("create", v1.CreateCategory)
			categoryGroup.PUT("update", v1.UpdateCategory)
			categoryGroup.DELETE("delete", v1.DeleteCategory)
		}
		fileGroup := managerGroup.Group("file")
		{

			fileGroup.DELETE("delete", v1.DeleteFile)
		}
	}
}
