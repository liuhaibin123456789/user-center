package global

import "time"

const (
	JWTSecret            = "123456"                         //JWT密钥
	TokenExpiresDuration = time.Hour * 12                   //Token过期时间段
	MvPicturePath        = "../static/picture/mvpicture/"   //电影头像存储目录
	UserAvatarPath       = "../static/picture/useravatar/"  //用户头像存储目录
	PerfPicturePath      = "../static/picture/perfpicture/" //影人头像储存目录
)
