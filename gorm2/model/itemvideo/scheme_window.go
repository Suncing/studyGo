package itemvideo

import "gorm.io/gorm"

/**
*   @author wangchenyang
*   @date 2022/8/29 15:16
*   @description:
 */
type SchemeWindow struct {
	gorm.Model
	SchemeId int `gorm:"type:int(10);comment:'方案id';not null" json:"schemeId"`
	WindowId int `gorm:"type:int(10);comment:'窗口id';not null" json:"windowId"`
	PosX     int `gorm:"type:int(0);comment:'x位置';not null" json:"posX"`
	PosY     int `gorm:"type:int(0);comment:'y位置';not null" json:"posY"`
	Width    int `gorm:"type:int(10);comment:'窗口宽度';not null" json:"width"`
	Height   int `gorm:"type:int(10);comment:'窗口高度';not null" json:"height"`
	//Window   Window `gorm:"foreignkey:WindowId"`
}
