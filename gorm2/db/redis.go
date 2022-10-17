package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm2/gorm2/config"
	"log"
)

var (
	CtxRedis = context.Background()
	Redis    *redis.Client
)

// NewRedis 返回redis实例
func NewRedis() (*redis.Client, error) {
	var redisClient *redis.Client
	//加载redis配置文件
	redisConfig := config.InitRedisConfig()
	if redisConfig != nil {
		//读取redis配置
		redisClient = redis.NewClient(&redis.Options{
			Network:  redisConfig.NetWork,
			Addr:     redisConfig.Addr + redisConfig.Port,
			Password: redisConfig.Password,
			DB:       redisConfig.Database,
		})
	} else {
		log.Fatalf("设置redis配置出错")
	}

	//心跳测试
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Println("Redis连接成功！")
	return redisClient, nil
}

func InitRedis() {
	newRedis, err := NewRedis()
	if err != nil {
		log.Fatalln("redis连接失败", err)
	}
	Redis = newRedis
}
