package config

import (
	"encoding/json"
	"log"
	"os"
)

type Mysql struct {
	DiverName string `json:"diver_name"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Addr      string `json:"addr"`
	DataBase  string `json:"data_base"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parse_time"`
	Loc       string `json:"loc"`
}

// InitMysqlConfig 初始化服务器配置
func InitMysqlConfig() *Mysql {
	file, err := os.Open("gorm2/config/mysql.json")
	if err != nil {
		log.Fatalln("读取redis配置出错", err)
	}
	//生成文件解析器
	decoder := json.NewDecoder(file)
	conf := Mysql{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
