package v1

import (
	"blog/model"
	"blog/model/request"
	"blog/model/response"
	"blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		response.FailWidthMessage("分类列表查询失败", c)
		return
	}
	response.OkWithData(categories, c)
}

func CreateCategory(c *gin.Context) {
	var param request.AddCategory
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthMessage("参数错误："+err.Error(), c)
		return
	}
	category := model.BlogCategory{
		Name: param.Name,
	}
	err = service.CreateCategory(&category)
	if err != nil {
		response.FailWidthMessage("创建失败", c)
		return
	}
	response.OkWithNil(c)
}
func UpdateCategory(c *gin.Context) {
	var category model.BlogCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWidthMessage("参数错误："+err.Error(), c)
		return
	}
	err = service.UpdateCategory(&category)
	if err != nil {
		response.FailWidthMessage("更新失败", c)
		return
	}
	response.OkWithNil(c)
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
		return
	}
	err = service.DeleteCategory(id)
	if err != nil {
		response.FailWidthMessage(err.Error(), c)
		return
	}
	response.OkWithNil(c)
}
