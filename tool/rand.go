package tool

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

//RandString 获取15字符的随机字符串
func RandString() string {
	str := []byte("qwertyuiopasdfghjklzxcvbnm")
	randString := make([]byte, 15)
	for _, i := range randString {
		rand.Seed(time.Now().UnixNano())
		randString[i] = str[rand.Intn(len(str))]
	}
	return string(randString)
}

//RandNum 生成10000-99999的五位随机数
func RandNum() (int, error) {
	randNum := rand.Intn(99999-10000) + 10000
	if randNum >= 10000 && randNum < 100000 {
		log.Println(randNum)
		return randNum, nil
	}
	return -1, errors.New("the rand code is wrong")
}

//RedisSetExp 随机码储存服务端redis
func RedisSetExp(v interface{}, phone string) error {
	statusCmd := RDB.Set(phone, v, time.Minute*5)
	_, err := statusCmd.Result()
	return err
}

//RedisGetExp 获取服务端redis里的随机码
func RedisGetExp(phone string) (string, error) {
	stringCmd := RDB.Get(phone)
	_, err := stringCmd.Result()
	if err != nil {
		return "", err
	}
	return stringCmd.Val(), err
}
