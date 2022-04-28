package model

//User 保存注册用户隐私数据
type User struct {
	UserID   int    `json:"user_id" form:"user_id" gorm:"primaryKey;autoIncrement"`
	UserName string `json:"user_name" form:"user_name" gorm:"type:varchar(40);not null"`
	Password string `json:"password" form:"password" gorm:"type:varchar(40);not null"`
	Phone    string `json:"phone" form:"phone" gorm:"type:varchar(11);unique;not null"`    //用户绑定的手机号码（唯一）
	Question string `json:"question" form:"question" gorm:"type:varchar(100);default:'无'"` //忘记密码时回答的问题
	Answer   string `json:"answer" form:"answer" gorm:"type:varchar(100);default:'无'"`     //忘记密码时的问题答案
}

func (User) TableName() string {
	return "user"
}

//UserSide 保存注册用户非隐私数据
type UserSide struct {
	UserId           int32  `json:"user_id" form:"user_id" gorm:"primaryKey;autoIncrement"`                     //为对应用户隐私表主键ID
	Phone            string `json:"phone" form:"phone" gorm:"type:varchar(11);unique;not null"`                 //当前登录用户的手机号
	Avatar           string `json:"avatar" form:"avatar" gorm:"type:varchar(200);default:'默认头像.jpg'"`           //用户头像绝对路径
	UserName         string `json:"user_name" form:"user_name" gorm:"type:varchar(40);not null"`                //用户的昵称,依据隐私表数据
	UserIntroduction string `json:"user_introduction" form:"user_introduction" gorm:"VARCHAR(1000);DEFAULT:''"` //用户的自我介绍
	UserSign         string `json:"user_sign" form:"user_sign" gorm:"type:VARCHAR(200);DEFAULT:''"`             //用户的签名档
	RegisterTime     string `json:"register_time" form:"register_time" gorm:"type:VARCHAR(25) ;NOT NULL"`       //用户的注册时间
}

func (UserSide) TableName() string {
	return "user_side"
}

//UserInfo 绑定数据，仅方便用于获取账户密码
type UserInfo struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}
