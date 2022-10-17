package config

import (
	"encoding/json"
	"log"
	"os"
)

/**
 * Redis 配置
 */
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

// InitRedisConfig 初始化服务器配置
func InitRedisConfig() *Redis {
	file, err := os.Open("gorm2/config/redis.json")
	if err != nil {
		log.Fatalln("读取redis配置出错", err)
	}
	//生成文件解析器
	decoder := json.NewDecoder(file)
	conf := Redis{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
