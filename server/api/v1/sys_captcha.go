package v1

import (
	"blog/global"
	"blog/global/response"
	resp "blog/model/response"
	"blog/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func Captcha(c *gin.Context) {
	captchaId := captcha.NewLen(global.GVA_CONFIG.Captcha.KeyLong)
	response.OkDetailed(resp.SysCaptchaResponse{
		CaptchaId: captchaId,
		PicPath:   "/api/base/captcha/" + captchaId + ".png",
	}, "验证码图片生成成功", c)
}
func CaptchaImg(c *gin.Context) {
	utils.GinCaptchaServeHTTP(c.Writer, c.Request)
}
