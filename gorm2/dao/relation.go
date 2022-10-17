package dao

import (
	"gorm2/gorm2/db"
	"gorm2/gorm2/model"
	"gorm2/gorm2/utils"
)

func SaveRelation() {
	p := model.People{
		Name: "wang",
		Company: model.Company{
			Name: "nova",
			ID:   1,
		},
	}
	db.DB.Create(&p)
}

func ReadRelation() {
	p := model.People{}
	db.DB.Preload("Company").First(&p)
	utils.ToString(p)
}

func PolymorphicCreate() {
	db.DB.Debug().Create(&model.Dog{Name: "dog1", Toy: model.Toy{Name: "toy1"}})
}

func HasManyCreate() {
	db.DB.Debug().Create(&model.Door{
		Key: []model.Key{
			{Number: "1"},
			{Number: "2"},
			{Number: "3"}},
	})
}

func SetupJoinCreate() {
	db.DB.SetupJoinTable(&model.Rich{}, "House", &model.RichAddress{})
}
