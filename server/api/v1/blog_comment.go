package v1

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"blog/model/response"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddComment(c *gin.Context) {
	var param request.AddComment
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthMessage("请填写评论内容", c)
		return
	}
	claims, exists := c.Get("claims")
	if !exists {
		response.FailWidthMessage("未查询到用户信息", c)
	}
	userClaims, _ := claims.(*model.CustomClaims)
	_, user := service.UserDetail(userClaims.Id)
	comment := model.BlogComment{
		ArticleId:  param.ArticleId,
		ParentId:   param.ParentId,
		Content:    param.Content,
		Username:   user.Username,
	}
	err = service.AddComment(&comment)
	if err != nil {
		response.FailWidthMessage("添加失败", c)
		return
	}
	response.OkWithData(comment, c)

}

func GetCommentList(c *gin.Context) {
	var param request.TreeComment
	err := c.ShouldBindJSON(&param)
	if err != nil {
		global.GVA_LOG.Error("参数错误", zap.Any("err", err))
		response.FailWidthMessage("参数错误", c)
		return
	}
	err, count, comments := service.GetCommentTree(param)
	if err != nil {
		response.FailWidthMessage("获取评论列表失败", c)
		return
	}
	result := response.PageResult{
		Total: count,
		List: comments,
		Page: param.PageNo,
		PageSize: param.PageSize,
	}
	response.OkWithData(result, c)

}
