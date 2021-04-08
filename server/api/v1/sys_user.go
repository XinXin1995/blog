package v1

import (
	"blog/global"
	"blog/global/response"
	"blog/middleware"
	"blog/model"
	"blog/model/request"
	resp "blog/model/response"
	"blog/service"
	"blog/utils"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func UserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetUserList(pageInfo)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("查询失败：%v", err.Error()), c)
	} else {
		result := &resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}
		response.OkDetailed(result, "查询成功", c)
	}
}

func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	var roles = make([]model.SysUserRole, len(R.Roles))
	Id := utils.CreateUuid()
	if len(R.Roles) > 0 {
		for index, role := range R.Roles {
			roles[index].UserId = Id
			roles[index].RoleId = role.Id
			roles[index].IsDefault = role.IsDefault
		}
	}
	user := &model.SysUser{
		Id:       Id,
		Username: R.Username,
		Password: R.Password,
		Nickname: R.Nickname,
		Avatar:   R.Avatar,
	}
	err := service.Register(*user, roles)
	if err != nil {
		response.FailWidthDetailed(response.ERROR, nil, fmt.Sprintf("%v", err.Error()), c)
	} else {
		response.OkDetailed(nil, "注册成功", c)
	}
}

func UserDetail(c *gin.Context) {
	Id := c.Query("id")
	if len(Id) == 0 {
		response.FailWidthMessage("请传入角色uuid", c)
		return
	}
	U := model.SysUser{
		Id: Id,
	}
	err, userReturn := service.UserDetail(U)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("查询失败：%v", err.Error()), c)
	} else {
		response.OkDetailed(userReturn, "操作成功", c)
	}
}

func Valid(c *gin.Context) {
	var V request.ValidStruct
	_ = c.ShouldBindJSON(&V)
	if captcha.VerifyString(V.CaptchaId, V.Captcha) {
		U := model.SysUser{Username: V.Username, Password: V.Password}
		if err, user := service.Login(U); err != nil {
			response.FailWidthMessage(fmt.Sprintf("用户名密码错误或v%", err.Error()), c)
		} else {
			tokenNext(c, user)
		}
	} else {
		response.FailWidthMessage("验证码错误", c)
	}
}

func UserBindRole(c *gin.Context) {
	var param request.UserBindRole
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("绑定角色失败：%v", err.Error()), c)
		return
	}
	userId := param.UserId
	var roles = make([]model.SysUserRole, len(param.Roles))
	if len(param.Roles) > 0 {
		for index, role := range param.Roles {
			roles[index].UserId = userId
			roles[index].RoleId = role.Id
			roles[index].IsDefault = role.IsDefault
		}
	}
	err = service.UserBindRole(userId, roles)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("绑定角色失败：%v", err.Error()), c)
	} else {
		response.OkWithNil(c)
	}

}

func tokenNext(c *gin.Context, user *resp.SysUserDetailResponse) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		Id:       user.Id,
		UserName: user.Username,
		Nickname: user.Nickname,
		Roles:    user.Roles,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "wcx",                                 //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWidthMessage("获取token失败", c)
	} else {
		response.OkWithData(resp.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}
