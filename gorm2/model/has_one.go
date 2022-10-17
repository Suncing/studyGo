package model

import "gorm.io/gorm"

type People2 struct {
	gorm.Model
	Name       string     `sql:"index"`
	CreditCard CreditCard `gorm:"foreignKey:UserName;references:name"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}
