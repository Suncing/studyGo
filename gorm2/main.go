package main

import (
	"gorm2/gorm2/dao"
	"gorm2/gorm2/db"
)

/**
*   @author wangchenyang
*   @date 2022/9/26 9:16
*   @description:
 */
func main() {
	// 初始化数据库
	db.InitRedis()
	// 创建服务
	db.InitMysql()
	dao.SetupJoinCreate()
}
