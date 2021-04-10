package v1

import (
	"blog/global"
	"blog/model"
	"blog/model/response"
	"blog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Upload(c *gin.Context) {
	t, _ := strconv.Atoi(c.PostForm("type"))
	var file model.SysFile
	_, header, err := c.Request.FormFile("file")
	var prefix string
	switch t {
	case 1:
		prefix = "user"
	case 2:
		prefix = "article"
	default:
		prefix = ""
	}
	if err != nil {
		global.GVA_LOG.Error("接受文件失败", zap.Any("err", err))
		response.FailWidthMessage("接受文件失败", c)
		return
	}
	err, file = service.UploadFile(header, prefix)

	if err != nil {
		global.GVA_LOG.Error("文件上传失败", zap.Any("err", err))
		response.FailWidthMessage("上传失败", c)
		return
	}
	response.OkWithData(file, c)
}

func DeleteFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWidthMessage("id错误", c)
		return
	}

	err = service.DeleteFile(id)
	if err != nil {
		response.FailWidthMessage("删除失败"+err.Error(), c)
		return
	}
	response.OkWithNil(c)

}
