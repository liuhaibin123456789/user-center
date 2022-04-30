package global

import "time"

const (
	JWTSecret            = "123456"                        //JWT密钥
	TokenExpiresDuration = time.Hour * 12                  //Token过期时间段
	UserAvatarPath       = "../static/picture/useravatar/" //用户头像存储目录
	Port                 = ":=8085"
)
