package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm2/gorm2/db"
	"gorm2/gorm2/model/itemvideo"
)

func DeleteAll() {
	user := itemvideo.User{}
	err := db.DB.Delete(&user).Error
	is := errors.Is(err, gorm.ErrMissingWhereClause)
	fmt.Println(is)
}
