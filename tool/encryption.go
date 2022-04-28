package tool

import (
	"crypto/md5"
	"fmt"
)

//Encrypt 加密字符串算法MD5
func Encrypt(oldStr string) string {
	//获取16位的md5
	newStr := md5.Sum([]byte(oldStr))
	return fmt.Sprintf("%x", newStr)
}
