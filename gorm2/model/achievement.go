package model

type Achievement struct {
	StudentId uint    `gorm:"type:int(20);comment:'学生id';unique_index" json:"studentId"`
	Math      float32 `gorm:"type:float(0);comment:'数学'" json:"math"`
	Chinese   float32 `gorm:"type:float(0);comment:'语文'" json:"chinese"`
	English   float32 `gorm:"type:float(0);comment:'英语'" json:"english"`
}
