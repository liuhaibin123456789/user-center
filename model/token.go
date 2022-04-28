package model

import "github.com/dgrijalva/jwt-go"

//Claims 用于获取token
type Claims struct {
	Phone              string `json:"phone" form:"phone"` //自定义字段
	jwt.StandardClaims        //官方字段
}
