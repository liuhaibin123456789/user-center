package tool

import (
	"github.com/go-redis/redis"
)

//RDB redis客户端
var RDB *redis.Client

//InitRedis 初始化redis连接
func InitRedis() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //todo 注意云服务端的配置需要改成相应的配置
		Password: "",               //todo 注意云服务端的配置需要改成相应的配置
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	RDB = rdb

	return nil
}
