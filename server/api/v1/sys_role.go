package v1

import (
	"blog/global/response"
	"blog/model"
	"blog/model/request"
	"blog/service"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func AddRole(c *gin.Context) {
	var R request.AddRoleStruct
	_ = c.ShouldBindJSON(&R)
	role := model.SysRole{
		Id:   utils.CreateUuid(),
		Name: R.Name,
		Desc: R.Desc,
	}
	err, _ := service.AddRole(role)
	if err != nil {
		response.FailWidthMessage("新建角色失败", c)
	} else {
		response.OkDetailed(nil, "新建角色成功", c)
	}
}

func RoleBindPage(c *gin.Context) {

}
