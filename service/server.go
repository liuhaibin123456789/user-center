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
func (s Service) Register(ctx context.Context, req *proto.ReqUser) (res *proto.ResToken, err error) {
	res = new(proto.ResToken)

	u := model.User{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone(),
		Question: req.GetQuestion(),
		Answer:   req.GetAnswer(),
	}
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
	log.Println(rand, err)
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
	//用户名为空，获取随机用户名GoString implements fmt.GoStringer and formats t to be printed in Go source code
	if u.UserName == "" {
		u.UserName = tool.RandString()
	}
	us := model.UserSide{UserName: u.UserName, Phone: u.Phone, RegisterTime: time.Now().Format("2006-01-02")}
	//添加用户隐私数据以及非隐私数据表
	err = dao.InsertUser(u, us)
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

}

func (s Service) GetQuestion(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {
	res = new(proto.ResUser)

	question, err := dao.SelectUserQuestion(req.GetPhone())
	if err != nil {
		res.OK = "fail"
		return res, err
	}
	res.OK = "success"
	//todo
	res.One = &proto.ReqUser_Question{Question: question}
	//将req的问题写入res的one里

	return res, err
}

func (s Service) GetSign(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}

func (s Service) GetAnswer(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}

func (s Service) CreateIntroduction(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}
func (s Service) CreateQuestion(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}

func (s Service) CreateSign(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}
func (s Service) CreateAnswer(ctx context.Context, req *proto.ReqUser) (res *proto.ResUser, err error) {

}
