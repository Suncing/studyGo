package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name        string `gorm:"type:varchar(30);default:'王晨阳';comment:'姓名'" json:"name"`
	Age         int    `gorm:"type:int(20);comment:'年龄'" json:"age"`
	Achievement Achievement
}
