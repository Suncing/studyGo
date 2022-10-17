package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm2/gorm2/db"
	"gorm2/gorm2/model/itemvideo"
)

// UpdateAll ,Save 会保存所有的字段，即使字段是零值
func UpdateAll() {
	user := itemvideo.User{}
	db.DB.First(&user)
	user.UserName = ""
	//user.Flag = 0
	db.DB.Save(&user)
}

func UpdateAllData() {
	err := db.DB.Model(&itemvideo.User{}).Update("user_name", "王晨阳").Error
	is := errors.Is(err, gorm.ErrMissingWhereClause)
	fmt.Println(is)
}

func UpdateOne() {
	db.DB.Model(&itemvideo.User{}).Where("id=?", 2).Update("pass_word", 0)
}

func UpdatesTest() {
	//err := db.DB.Model(&itemvideo.User{}).Where("id=?", 2).Updates(itemvideo.User{UserName: "", Flag: 0}).Error
	err := db.DB.Model(&itemvideo.User{}).Where("id=?", 2).Updates(map[string]interface{}{"user_name": "", "flag": 0}).Error
	is := errors.Is(err, gorm.ErrMissingWhereClause)
	fmt.Println(is)
}
