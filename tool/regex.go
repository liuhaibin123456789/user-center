package tool

import (
	"errors"
	"regexp"
)

//RegexPhone 匹配手机号的格式，11位数字
func RegexPhone(phone string) error {
	if phone == "" {
		return errors.New("手机号为空")
	}
	regexStr := "[0-9]{11}"
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}
	if !regex.MatchString(phone) {
		return errors.New("the pattern of phone is wrong." +
			"the phone must be 11 numbers except others")
	}
	return err
}

//RegexUserNameAndPwd 匹配用户名及密码的格式，
func RegexUserNameAndPwd(checkingStr string) error {
	if checkingStr == "" {
		return errors.New("账号或密码为空")
	}
	regexStr := "[a-zA-Z0-9]{8,16}"
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}
	if !regex.MatchString(checkingStr) {
		return errors.New("the pattern of userName or password is wrong," +
			"user_name must be Uppercase , lowercase letters or numbers range from 8 to 16 digits")
	}
	return err
}

//RegexChinese 校验是否为中文
func RegexChinese(str string) error {
	reg := "[\u4e00-\u9fa5]+"
	regex, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	if !regex.MatchString(str) {
		return errors.New("the string is not chinese")
	}
	return err
}

//RegexSearch 匹配搜索内容
func RegexSearch(str string) error {
	reg := "[a-zA-Z0-9_\u4e00-\u9fa5]+"
	regex, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	if !regex.MatchString(str) {
		return errors.New("请输入中英文数字或下划线")
	}
	return err
}
