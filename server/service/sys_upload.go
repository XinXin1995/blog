package service

import (
	"blog/global"
	"blog/model"
	ossutil "blog/utils/upload"
	"errors"
	"mime/multipart"
	"strings"
)

func UploadFile(header *multipart.FileHeader, prefix string) (err error, file model.SysFile) {
	oss := &ossutil.TencentCos{}

	filePath, key, uploadErr := oss.UploadFile(header, prefix)

	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := model.SysFile{
		Url:    filePath,
		Name:   header.Filename,
		Prefix: prefix,
		Tag:    s[len(s)-1],
		Key:    key,
	}
	err = upload(&f)
	return err, f
}

func DeleteFile(id int) (err error) {
	var fileFromDb model.SysFile
	err, fileFromDb = FindFile(uint(id))
	oss := &ossutil.TencentCos{}
	if err = oss.DeleteFile(fileFromDb); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", fileFromDb.ID).Unscoped().Delete(&fileFromDb).Error
	return err
}

func FindFile(id uint) (error, model.SysFile) {
	var file model.SysFile
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

func upload(file *model.SysFile) error {
	return global.GVA_DB.Create(&file).Error
}
