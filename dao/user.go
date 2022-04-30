package dao

import (
	"fmt"
	"user-center/global"
	"user-center/model"
	"user-center/tool"
)

//InsertUser 同时插入隐私表、非隐私表数据
func InsertUser(user1 model.User, user2 model.UserSide) error {
	tx := tool.GDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Model(&model.User{}).Create(&user1).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.UserSide{}).Create(&user2).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateUserPwd(phone, userPwd string) error {
	return tool.GDb.Model(&model.User{}).Where("phone=?", phone).Update("password", userPwd).Error
}

func SelectUserPwd(phone string) (string, error) {
	var pwd string
	if err := tool.GDb.Model(&model.User{}).Select("password").Where("phone=?", phone).Find(&pwd).Error; err != nil {
		return "", err
	}
	return pwd, nil
}

func SelectUserPhone(phone string) (string, error) {
	var un string
	if err := tool.GDb.Model(&model.UserSide{}).Select("user_name").Where("phone=?", phone).Find(&un).Error; err != nil {
		return "", err
	}
	return un, nil
}

func SelectUserAnswer(phone string) (string, error) {
	var answer string
	if err := tool.GDb.Model(&model.User{}).Select("answer").Where("phone=?", phone).Find(&answer).Error; err != nil {
		return "", err
	}
	return answer, nil

}

func SelectUserQuestion(phone string) (string, error) {
	var question string
	if err := tool.GDb.Model(&model.User{}).Select("question").Where("phone=?", phone).Find(&question).Error; err != nil {
		return "", err
	}
	return question, nil
}

func SelectUserSign(phone string) (string, error) {
	var sign string
	if err := tool.GDb.Model(&model.User{}).Select("question").Where("phone=?", phone).Find(&sign).Error; err != nil {
		return "", err
	}
	return sign, nil
}

func SelectUserIntroduction(phone string) (string, error) {
	var introduction string
	if err := tool.GDb.Model(&model.UserSide{}).Select("user_introduction").Where("phone=?", phone).Find(&introduction).Error; err != nil {
		return "", err
	}
	return introduction, nil
}

func InsertUserIntroduction(phone, introduction string) error {
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("user_introduction", introduction).Error
}

func InsertUserQuestion(phone, question string) error {
	return tool.GDb.Model(&model.User{}).Where("phone=?", phone).Update("question", question).Error
}

func InsertUserSign(phone, sign string) error {
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("user_sign", sign).Error
}

func InsertUserAnswer(phone, answer string) error {
	return tool.GDb.Model(&model.User{}).Where("phone=?", phone).Update("answer", answer).Error
}

func InsertUserAvatar(phone, filename string) error {
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("avatar", filename).Error
}

func SelectUserSide(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	if err := tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Find(&us).Error; err != nil {
		return model.UserSide{}, err
	}
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, nil
}

func SelectOInfo(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	if err := tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Find(&us).Error; err != nil {
		return model.UserSide{}, err
	}
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, nil
}
