package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm2/gorm2/config"
	"gorm2/gorm2/model"
	"gorm2/gorm2/model/itemvideo"
	"log"
	"strconv"
	"time"
)

var (
	DB *gorm.DB
)

func NewMysql() string {
	mysqlConfig := config.InitMysqlConfig()
	if mysqlConfig == nil {
		log.Fatalf("设置mysql配置出错")
	}
	//driverName := mysqlConfig.DiverName
	connInfo := mysqlConfig.User + ":" +
		mysqlConfig.Password + "@(" +
		mysqlConfig.Addr + ")/" +
		mysqlConfig.DataBase + "?charset=" +
		mysqlConfig.Charset + "&parseTime=" +
		strconv.FormatBool(mysqlConfig.ParseTime) + "&loc=" +
		mysqlConfig.Loc
	return connInfo
}

func createDB() *gorm.DB {
	newMysql := NewMysql()

	db, err := gorm.Open(mysql.Open(newMysql), &gorm.Config{})

	if err != nil {
		log.Fatalln("mysql  conn err:", err)
	}
	log.Println("mysql数据库连接成功！")
	return db
}

func InitMysql() {
	db := createDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("初始化数据库出错", err)
	}
	// 开启日志
	db.Logger.LogMode(1)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// db.SingularTable(true) // 禁用复数形式
	// 设置表名称前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "video_" + defaultTableName
	//}
	DB = db
	// 注册表结构
	initModel(db)
}

func initModel(db *gorm.DB) {
	err := db.AutoMigrate(
		&itemvideo.User{},
		&itemvideo.Window{},
		&itemvideo.SchemeWindow{},
		&model.Student{},
		&model.Achievement{},
		&model.People{},
		&model.Company{},
		//&model.CreditCard{},
		//&model.People2{},
		&model.Cat{},
		&model.Dog{},
		&model.Toy{},
		&model.Door{},
		&model.Key{},
		&model.Language{},
		&model.Scholar{},
		&model.Pupil{},
		&model.Profile{},
		&model.Rich{},
		&model.House{},
	)
	if err != nil {
		log.Fatalln("迁移出错", err)
	}
}
