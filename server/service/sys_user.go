package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"blog/utils"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//用户列表
func GetUserList(p request.PageInfo) (err error, list []model.SysUser, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.PageNo - 1)
	db := global.GVA_DB
	err = db.Where("username like ?", "%"+p.Keyword+"%").Limit(limit).Offset(offset).Find(&list).Count(&total).Error
	return
}

//用户注册（添加用户）
func Register(u *model.SysUser) (err error) {
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("用户名已存在", zap.Any("err", err))
		return errors.New("用户名已存在")
	}
	u.Password = utils.MD5V([]byte(u.Password))

	return global.GVA_DB.Create(&u).Error
}

// 登录
func Login(u model.SysUser) (error, model.SysUser) {
	u.Password = utils.MD5V([]byte(u.Password))
	err := global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名或密码不正确"), model.SysUser{}
	}
	return err, u
}

// 获取用户信息详情
func UserDetail(id uint) (err error, userReturn model.SysUser) {
	if errors.Is(global.GVA_DB.Where("id = ?", id).First(&userReturn).Error, gorm.ErrRecordNotFound) {
		return
	}
	return
}
