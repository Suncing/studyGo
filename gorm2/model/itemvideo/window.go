package itemvideo

import "gorm.io/gorm"

type Window struct {
	gorm.Model
	UserId     uint   `gorm:"type:int(10);comment:'用户id';not null" json:"userId"`
	WindowName string `gorm:"type:varchar(30);comment:'窗口名称';not null" json:"windowName"`
	PosX       int    `gorm:"type:int(0);comment:'x位置';not null" json:"posX"`
	PosY       int    `gorm:"type:int(0);comment:'y位置';not null" json:"posY"`
	Width      int    `gorm:"type:int(10);comment:'窗口宽度';not null" json:"width"`
	Height     int    `gorm:"type:int(10);comment:'窗口高度';not null" json:"height"`
}
