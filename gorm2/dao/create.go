package dao

import (
	"fmt"
	"gorm.io/gorm/clause"
	"gorm2/gorm2/db"
	"gorm2/gorm2/model"
	"gorm2/gorm2/model/itemvideo"
	"log"
)

// Insert 单插
func Insert() {
	user := itemvideo.User{
		UserName:     "naruto",
		PassWord:     "1234",
		HeadPortrait: "123.jpg",
		NickName:     "鸣人",
		Email:        "980885033@qq.com",
		Tel:          "13571915380",
		QuestionA:    "西安",
		QuestionB:    "黑黎",
		QuestionC:    "刘亦菲",
	}
	result := db.DB.Create(&user) // 通过数据的指针来创建
	// 返回插入数据的主键
	// 返回插入数据的主键
	log.Println(user.ID)
	// 返回 error
	log.Println(result.Error)
	// 返回插入记录的条数
	log.Println(result.RowsAffected)
	//获取插入记录的Id
	var id uint
	db.DB.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
}

// InsertSelect 用选定字段的来创建
func InsertSelect() {
	user := itemvideo.User{
		UserName:     "naruto",
		PassWord:     "1234",
		HeadPortrait: "123.jpg",
		NickName:     "鸣人",
		Email:        "980885033@qq.com",
		Tel:          "13571915380",
		QuestionA:    "西安",
		QuestionB:    "黑黎",
		QuestionC:    "刘亦菲",
	}
	db.DB.Select("user_name", "pass_word").Create(&user) // 通过数据的指针来创建
}

// InsertOmit 建时排除选定字段
func InsertOmit() {
	user := itemvideo.User{
		UserName:     "naruto",
		PassWord:     "1234",
		HeadPortrait: "123.jpg",
		NickName:     "鸣人",
		Email:        "980885033@qq.com",
		Tel:          "13571915380",
		QuestionA:    "西安",
		QuestionB:    "黑黎",
		QuestionC:    "刘亦菲",
	}
	db.DB.Omit("user_name", "pass_word").Create(&user) // 通过数据的指针来创建
}

// InsertBatch 批量创建
func InsertBatch() {
	users := []itemvideo.User{
		{UserName: "naruto",
			PassWord:     "1234",
			HeadPortrait: "123.jpg",
			NickName:     "鸣人",
			Email:        "980885033@qq.com",
			Tel:          "13571915380",
			QuestionA:    "西安",
			QuestionB:    "黑黎",
			QuestionC:    "刘亦菲"},
		{UserName: "lisa",
			PassWord:     "1234",
			HeadPortrait: "123.jpg",
			NickName:     "lisa",
			Email:        "980885033@qq.com",
			Tel:          "13571915380",
			QuestionA:    "西安",
			QuestionB:    "黑黎",
			QuestionC:    "刘亦菲"},
	}
	db.DB.Create(&users) // 通过数据的指针来创建
	for _, user := range users {
		fmt.Println(user.ID)
	}
}

// InsertRelation 关联创建
func InsertRelation() {
	student := model.Student{
		Name: "wangcy",
		Age:  26,
		Achievement: model.Achievement{
			//成绩的学生外键id会自动生成插入
			Math:    98,
			Chinese: 99.5,
			English: 90,
		},
	}
	db.DB.Create(&student)
}

// InsertRelationSelect 跳过关联保存
func InsertRelationSelect() {
	student := model.Student{
		Name: "wangcy",
		Age:  26,
		Achievement: model.Achievement{
			//成绩的学生外键id会自动生成插入
			Math:    98,
			Chinese: 99.5,
			English: 90,
		},
	}
	db.DB.Select("Achievement").Create(&student)
}

// InsertRelationOmit 跳过关联保存
func InsertRelationOmit() {
	student := model.Student{
		Name: "wangcy",
		Age:  26,
		Achievement: model.Achievement{
			//成绩的学生外键id会自动生成插入
			Math:    98,
			Chinese: 99.5,
			English: 90,
		},
	}
	db.DB.Omit("Achievement").Create(&student)
	// 跳过所有关联
	db.DB.Omit(clause.Associations).Create(&student)
}

// InsertZero 注意 像 0、''、false 等零值，不会将这些字段定义的默认值保存到数据库。您需要使用指针类型或 Scanner/Valuer 来避免这个问题，例如：
// InsertZero 当数据库字段有默认值时，插入0值方法
func InsertZero() {
	//flag := 0
	user := itemvideo.User{
		UserName:     "naruto",
		PassWord:     "1234",
		HeadPortrait: "123.jpg",
		NickName:     "鸣人",
		Email:        "980885033@qq.com",
		Tel:          "13571915380",
		QuestionA:    "西安",
		QuestionB:    "黑黎",
		QuestionC:    "刘亦菲",
		//Flag:         &flag,
	}
	db.DB.Save(&user)
}
