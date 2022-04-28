package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-center/global"
	"user-center/model"
)

//CreateToken 获取token
func CreateToken(phone string) (string, error) {
	claim := model.Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.TokenExpiresDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cold bin & tao rui",
		},
	}
	//使用hs256算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(global.JWTSecret))
}

//ParseToken 解析浏览器端的token字符串，claims不为nil
func ParseToken(tokenString string) (*model.Claims, error) {
	//解析tokenString，拿到token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&model.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(global.JWTSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
