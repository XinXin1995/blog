package utils

import (
	"blog/global"
	"blog/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWidthMessage("验证码获取失败", c)
	} else {
		response.OkDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

//func GinCaptchaServeHTTP(w http.ResponseWriter, r *http.Request) {
//	dir, file := path.Split(r.URL.Path)
//	ext := path.Ext(file)
//	id := file[:len(file)-len(ext)]
//
//	if ext == "" && id == "" {
//		http.NotFound(w, r)
//		return
//	}
//	fmt.Println("reload : " + r.FormValue("reload"))
//	if r.FormValue("reload") != "" {
//		captcha.Reload(id)
//	}
//	lang := strings.ToLower(r.FormValue("lang"))
//	download := path.Base(dir) == "download"
//	if Serve(w, r, id, ext, lang, download, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.ImgHeight) == captcha.ErrNotFound {
//		http.NotFound(w, r)
//	}
//}
//
//func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
//	w.Header().Set("Cache-control", "no-cache,no-store,must-revalidate")
//	w.Header().Set("pragma", "no-cache")
//	w.Header().Set("Expires", "0")
//	var contents bytes.Buffer
//	switch ext {
//	case ".png":
//		w.Header().Set("Content-Type", "image/png")
//		_ = captcha.WriteImage(&contents, id, width, height)
//	case ".wav":
//		w.Header().Set("Content-Type", "audio/x-wav")
//		_ = captcha.WriteAudio(&contents, id, lang)
//	default:
//		return captcha.ErrNotFound
//	}
//	if download {
//		w.Header().Set("Content-Type", "application/octet-stream")
//	}
//	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(contents.Bytes()))
//	return nil
//}
