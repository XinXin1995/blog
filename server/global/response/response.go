package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, data, msg})
}

func OkWithNil(c *gin.Context) {
	Result(SUCCESS, nil, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func FailWidthDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func FailWidthMessage(m string, c *gin.Context) {
	Result(ERROR, nil, m, c)
}

func FailWidthError(msg string, err error, c *gin.Context) {
	Result(ERROR, nil, fmt.Sprintf("%s:%v", msg, err.Error()), c)

}
