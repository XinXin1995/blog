package request

import (
	"blog/model/response"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserName string
	Id       string
	Nickname string
	Roles    []response.SysUserRole
	jwt.StandardClaims
}
