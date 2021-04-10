package v1

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"blog/model/response"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func AddArticle(c *gin.Context) {
	var articleParam request.AddArticleParam
	if err := c.ShouldBindJSON(&articleParam); err != nil {
		response.FailWidthMessage(err.Error(), c)
		return
	}
	article := &model.BlogArticle{
		Title:      articleParam.Title,
		Content:    articleParam.Content,
		CategoryId: articleParam.Category,
		Thumb:      articleParam.Thumb,
	}
	err := service.AddArticle(article, articleParam.Tags)
	if err != nil {
		global.GVA_LOG.Error("新建文章失败", zap.Any("err", err.Error()))
		response.FailWidthMessage(err.Error(), c)
	} else {
		response.OkWithData(article, c)
	}
}

func GetArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
	} else {
		articleResp, err := service.GetArticle(articleId)
		if err != nil {
			response.FailWidthMessage(err.Error(), c)
		} else {
			response.OkWithData(articleResp, c)
		}
	}
}

func GetArticleList(c *gin.Context) {
	var param request.ArticleList
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthMessage(err.Error(), c)
		return
	}
	err, count, list := service.GetArticleList(&param)
	if err != nil {
		response.FailWidthMessage("查询失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:     list,
		Total:    count,
		PageSize: param.PageSize,
		Page:     param.PageNo,
	}, c)
}

func UpdateArticle(c *gin.Context) {
	var param request.UpdateArticleParam
	if err := c.ShouldBindJSON(&param); err != nil {
		response.FailWidthMessage(err.Error(), c)
		return
	}

	err := service.UpdateArticle(&param)
	if err != nil {
		global.GVA_LOG.Error("更新文章失败", zap.Any("err", err))
		response.FailWidthMessage("更新失败", c)
		return
	}
	response.OkWithNil(c)
}

func DelArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
		return
	}
	if err = service.DelArticle(id); err != nil {
		response.FailWidthMessage("删除失败", c)
		return
	}
	response.OkWithNil(c)
}

func LikeArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
		return
	}
	if err := service.LikeArticle(id); err != nil {
		response.FailWidthMessage("更新失败", c)
	} else {
		response.OkWithNil(c)
	}

}
