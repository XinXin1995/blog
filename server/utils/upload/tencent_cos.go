package upload

import (
	"blog/global"
	"blog/model"
	"context"
	"errors"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type TencentCos struct{}

func (t *TencentCos) UploadFile(file *multipart.FileHeader, prefix string) (string, string, error) {
	client := NewClient()
	f, openErr := file.Open()
	if openErr != nil {
		global.GVA_LOG.Error("function file.Open() Filed", zap.Any("err", openErr.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openErr.Error())
	}
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), prefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.GVA_CONFIG.TencentCOS.BaseURL + "/" + prefix + "/" + fileKey, fileKey, nil
}

func (t *TencentCos) DeleteFile(file model.SysFile) error {
	client := NewClient()
	name := file.Prefix + "/" + file.Key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.GVA_CONFIG.TencentCOS.Bucket + ".cos." + global.GVA_CONFIG.TencentCOS.Region + ".myqcloud.com")
	baseUrl := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.GVA_CONFIG.TencentCOS.SecretID,
			SecretKey: global.GVA_CONFIG.TencentCOS.SecretKey,
		},
	})
	return client
}
