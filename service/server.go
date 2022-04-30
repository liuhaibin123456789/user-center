package service

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"
	"user-center/dao"
	"user-center/model"
	"user-center/proto"
	"user-center/tool"
)

type Service struct {
	proto.UserCenterServer
	//内嵌接口，可以保障方法实现完成度
}

// Login 登录服务
func (s Service) Login(ctx context.Context, req *proto.ReqUser) (res *proto.ResToken, err error) {
	res = new(proto.ResToken)

	phone := req.GetPhone()
	password := req.GetPassword()

	userPwd, err := dao.SelectUserPwd(phone)
	if err != nil {
		return res, err
	}
	//加密算法
	password = tool.Encrypt(password)
	if userPwd != password {
		res.OK = "fail"
		return res, errors.New("the user`s password is wrong")
	}
	//获取token
	token, err := tool.CreateToken(phone)
	res.OK = "success"
	res.Token = token
	return res, nil
}

//Register 注册服务
func (s Service) Register(ctx context.Context, req *proto.UserInfo) (res *proto.ResToken, err error) {
	res = new(proto.ResToken)

	u := model.User{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone(),
		Question: req.GetQuestion(),
		Answer:   req.GetAnswer(),
	}
	if u.UserName == "" {
		u.UserName = tool.RandString()
	}
	//log.Println(u)
	//用户信息校验格式正则校验
	err = tool.RegexPhone(u.Phone)
	if err != nil {
		return res, err
	}
	err = tool.RegexUserNameAndPwd(u.UserName)
	if err != nil {
		return res, err
	}
	err = tool.RegexUserNameAndPwd(u.Password)
	if err != nil {
		return res, err
	}
	//检测随机码
	rand, err := tool.RedisGetExp(u.Phone)
	log.Println(rand)
	if err != nil {
		return res, errors.New("the rand code is out of date or you should put into right rand code")
	}
	//手机号存在的话，直接登录
	phone, err := dao.SelectUserPhone(u.Phone)
	if err == nil && u.Phone == phone {
		return res, err
	}
	//密码加密
	u.Password = tool.Encrypt(u.Password)
	us := model.UserSide{
		UserName:     u.UserName,
		Phone:        u.Phone,
		RegisterTime: time.Now().Format("2006-01-02"),
	}
	//添加用户隐私数据以及非隐私数据表
	err = dao.InsertUser(u, us)
	if err != nil {
		return res, err
	}
	res.Token, _ = tool.CreateToken(u.Phone)
	return res, err
}

func (s Service) GetCode(ctx context.Context, req *proto.ReqUser) (res *proto.ResCode, err error) {
	res = new(proto.ResCode)
	err = tool.RegexPhone(req.GetPhone())
	if err != nil {
		res.Code = "-1"
		res.OK = "failed"
		return res, err
	}
	// 获取网站随机码充当验证码
	randNum, err := tool.RandNum()
	if err != nil {
		res.Code = "-1"
		res.OK = "fail"
		return res, err
	}
	err = tool.RedisSetExp(randNum, req.GetPhone())
	if err != nil {
		res.Code = "-1"
		res.OK = "fail"
		return res, err
	}
	res.Code = strconv.Itoa(randNum)
	res.OK = "success"
	return res, err
}

func (s Service) GetUser(ctx context.Context, req *proto.ReqUser) (res *proto.UserSide, err error) {
	res = new(proto.UserSide)
	us, err := dao.SelectUserSide(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.UserId = us.UserId
	res.Avatar = us.Avatar
	res.Phone = us.Phone
	res.UserName = us.UserName
	res.RegisterTime = us.RegisterTime
	res.UserSign = us.UserSign
	res.UserIntroduction = us.UserIntroduction
	res.OK = "success"
	return res, err
}

func (s Service) GetIntroduction(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {
	res = new(proto.ResUser)

	introduction, err := dao.SelectUserIntroduction(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	res.One = &proto.ResUser_UserIntroduction{UserIntroduction: introduction}
	log.Println("(s Service) GetIntroduction: ", res)

	return res, err
}

func (s Service) GetQuestion(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {
	res = new(proto.ResUser)

	question, err := dao.SelectUserQuestion(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	res.One = &proto.ResUser_Question{Question: question}
	//todo viper日志库
	log.Println("(s Service) GetQuestion: ", res)
	return res, err
}

func (s Service) GetSign(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {
	res = new(proto.ResUser)

	question, err := dao.SelectUserSign(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	res.One = &proto.ResUser_Question{Question: question}
	log.Println("(s Service) GetSign: ", res)

	return res, err
}

func (s Service) GetAnswer(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {
	res = new(proto.ResUser)

	answer, err := dao.SelectUserAnswer(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	res.One = &proto.ResUser_Answer{Answer: answer}
	log.Println("(s Service) GetAnswer: ", res)

	return res, err
}

func (s Service) CreateIntroduction(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)
	//校验数据格式
	if len(req.GetIntroduction()) > 1000 {
		res.OK = "fail"
		return res, errors.New("the introduction is too long, please to write <1000 bytes")
	}

	err = dao.InsertUserIntroduction(req.GetPhone(), req.GetIntroduction())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) CreateIntroduction: ", res)

	return res, err
}
func (s Service) CreateQuestion(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)

	err = dao.InsertUserQuestion(req.GetPhone(), req.GetQuestion())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) CreateQuestion: ", res)

	return res, err
}

func (s Service) CreateSign(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)
	//校验数据格式
	if len(req.GetSign()) > 200 {
		res.OK = "fail"
		return res, errors.New("the sign is too long, please to write <200 bytes")
	}
	err = dao.InsertUserSign(req.GetPhone(), req.GetSign())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) CreateQuestion: ", res)

	return res, err
}
func (s Service) CreateAnswer(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)

	err = dao.InsertUserAnswer(req.GetPhone(), req.GetAnswer())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) CreateAnswer: ", res)

	return res, err
}

func (s Service) CreateAvatar(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)
	//校验数据格式
	if len(req.GetAvatar()) > 200 {
		res.OK = "fail"
		return res, errors.New("the filename is too long, please to write <200 bytes")
	}

	err = dao.InsertUserAvatar(req.GetPhone(), req.GetAvatar())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) CreateAvatar: ", res)

	return res, err
}

func (s Service) UpdatePwd(ctx context.Context, req *proto.ReqUser) (res *proto.Res, err error) {
	res = new(proto.Res)
	//校验密码格式
	err = tool.RegexUserNameAndPwd(req.GetPassword())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	err = tool.RegexUserNameAndPwd(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	//验证密码是否正确
	pwd, err := dao.SelectUserPwd(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	//加密算法
	oldPassword := tool.Encrypt(req.GetOldPassword())
	if pwd != oldPassword {
		res.OK = "fail"
		return res, errors.New("the password is wrong")
	}
	//加密算法
	newPassword := tool.Encrypt(req.GetPassword())
	err = dao.UpdateUserPwd(req.GetPhone(), newPassword)
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	log.Println("(s Service) UpdatePwd: ", res)

	return res, err
}
