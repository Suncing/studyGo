package model

import "gorm.io/gorm"

type Door struct {
	gorm.Model
	Key []Key
}

type Key struct {
	gorm.Model
	Number string
	DoorID uint
}
