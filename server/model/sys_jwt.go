package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id       uint
	UserName string
	Role     int
	jwt.StandardClaims
}
