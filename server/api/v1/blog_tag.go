package v1

import (
	"blog/model"
	"blog/model/request"
	"blog/model/response"
	"blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTags(c *gin.Context) {
	tags, err := service.GetTags()
	if err != nil {
		response.FailWidthMessage("标签列表查询失败", c)
		return
	}
	response.OkWithData(tags, c)
}

func CreateTag(c *gin.Context) {
	var param request.AddTag
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthMessage("参数错误："+err.Error(), c)
		return
	}
	tag := model.BlogTag{
		Name:  param.Name,
		Color: param.Color,
	}
	err = service.CreateTag(&tag)
	if err != nil {
		response.FailWidthMessage("更新失败", c)
		return
	}
	response.OkWithNil(c)
}
func UpdateTag(c *gin.Context) {
	var tag model.BlogTag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWidthMessage("参数错误："+err.Error(), c)
		return
	}
	err = service.UpdateTag(&tag)
	if err != nil {
		response.FailWidthMessage("更新失败", c)
		return
	}
	response.OkWithNil(c)
}

func DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
		return
	}
	err = service.DeleteTag(id)
	if err != nil {
		response.FailWidthMessage("删除失败", c)
		return
	}
	response.OkWithNil(c)
}
