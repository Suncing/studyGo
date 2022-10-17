package model

import (
	"gorm.io/gorm"
	"time"
)

type Scholar struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name     string
	Scholars []*Scholar `gorm:"many2many:user_languages;"`
}

type Pupil struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:PupilReferID222222;References:PupilRefer;joinReferences:ProfileRefer"`
	Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
	gorm.Model
	Name       string
	PupilRefer uint `gorm:"index:,unique"`
}

type Rich struct {
	ID    int
	Name  string
	House []House `gorm:"many2many:rich_house;"`
}

type House struct {
	ID   uint
	Name string
}

type RichAddress struct {
	RichID    int
	HouseID   int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
