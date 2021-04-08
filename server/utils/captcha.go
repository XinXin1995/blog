package utils

import (
	"blog/global"
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"net/http"
	"path"
	"strings"
	"time"
)

func GinCaptchaServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]

	if ext == "" && id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.ImgHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-control", "no-cache,no-store,must-revalidate")
	w.Header().Set("pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var contents bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&contents, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&contents, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(contents.Bytes()))
	return nil
}
