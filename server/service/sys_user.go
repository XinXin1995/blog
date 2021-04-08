package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"blog/model/response"
	"blog/utils"
	"errors"
	"fmt"
)

//用户列表
func GetUserList(p request.PageInfo) (err error, list []model.SysUser, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	db := global.GVA_DB
	err = db.Where("username like ?", "%"+p.Keyword+"%").Limit(limit, offset).Find(&list)
	total, err = db.Where("username like ?", "%"+p.Keyword+"%").Count(&model.SysUser{})
	return
}

//用户注册（添加用户）
func Register(u model.SysUser, roles []model.SysUserRole) (err error) {
	var user model.SysUser
	has, err := global.GVA_DB.Where("username = ?", u.Username).Get(&user)
	if has {
		return errors.New("用户名已注册")
	} else {
		u.Password = utils.MD5V([]byte(u.Password))
		session := global.GVA_DB.NewSession()
		defer session.Close()
		err = session.Begin()
		_, err = session.Insert(&u)
		if err != nil {
			session.Rollback()
			return
		}
		if len(roles) > 0 {
			for _, role := range roles {
				_, err = session.Insert(&role)
				if err != nil {
					session.Rollback()
					return
				}
			}
		}
		err = session.Commit()
		if err != nil {
			return
		} else {

		}
	}
	return err
}

// 用户绑定角色
func UserBindRole(userId string, roles []model.SysUserRole) (err error) {
	session := global.GVA_DB.NewSession()

	defer session.Close()

	err = session.Begin()

	if _, err = session.Exec("Delete From sys_user_role WHERE user_id = ?", userId); err != nil {
		session.Rollback()
		return
	}

	if len(roles) > 0 {
		for _, role := range roles {
			_, err = session.Insert(&role)
			if err != nil {
				session.Rollback()
				return
			}
		}
	}
	err = session.Commit()
	return
}

// 登录
func Login(u model.SysUser) (error, *response.SysUserDetailResponse) {
	u.Password = utils.MD5V([]byte(u.Password))
	has, err := global.GVA_DB.Get(&u)
	if !has {
		err = errors.New("用户不存在")
	}
	err, userReturn := UserDetail(u)
	return err, &userReturn
}

// 获取用户信息详情
func UserDetail(u model.SysUser) (err error, userReturn response.SysUserDetailResponse) {
	has, err := global.GVA_DB.Get(&u)
	if !has {
		err = errors.New("未查询到该用户信息")
		return
	}
	rolesSql := "SELECT r.id, r.name, ur.is_default FROM sys_user_role ur LEFT JOIN sys_role r ON ur.role_id = r.id WHERE ur.user_id = '%s'"
	var roles []response.SysUserRole
	err = global.GVA_DB.SQL(fmt.Sprintf(rolesSql, u.Id)).Find(&roles)
	userReturn.GetFomSysUser(&u)
	userReturn.Roles = roles
	return
}
