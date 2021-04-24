package v1

import (
	"blog/global"
	"blog/middleware"
	"blog/model"
	"blog/model/request"
	resp "blog/model/response"
	"blog/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func UserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetUserList(pageInfo)
	if err != nil {
		resp.FailWidthMessage(fmt.Sprintf("查询失败：%v", err.Error()), c)
	} else {
		result := &resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.PageNo,
			PageSize: pageInfo.PageSize,
		}
		resp.OkDetailed(result, "查询成功", c)
	}
}

func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := model.SysUser{
		Username: R.Username,
		Password: R.Password,
		Email:   R.Email,
	}
	err := service.Register(&user)
	if err != nil {
		resp.FailWidthDetailed(nil, fmt.Sprintf("%v", err.Error()), c)
	} else {
		resp.OkDetailed(nil, "注册成功", c)
	}
}

func UserDetail(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*model.CustomClaims)
	err, u := service.UserDetail(waitUse.Id)
	if err != nil {
		resp.FailWidthMessage(fmt.Sprintf("查询失败：%v", err.Error()), c)
	} else {
		resp.OkDetailed(u, "操作成功", c)
	}
}

func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		u := model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(u); err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			resp.FailWidthMessage("用户名不存在或者密码错误", c)
		} else {
			tokenNext(c, &user)
		}
	} else {
		resp.FailWidthMessage("验证码错误", c)
	}
}
func LoginForVisitor (c *gin.Context)  {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	u := model.SysUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(u); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
		resp.FailWidthMessage("用户名不存在或者密码错误", c)
	} else {
		tokenNext(c, &user)
	}
}

func tokenNext(c *gin.Context, user *model.SysUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := model.CustomClaims{
		Id:       user.ID,
		UserName: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "wcx",                                 //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		resp.FailWidthMessage("获取token失败", c)
	} else {
		resp.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}
