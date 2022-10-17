package itemvideo

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"type:varchar(20);comment:'用户名';unique_index" json:"userName"`
	PassWord     string `gorm:"type:varchar(35);comment:'密码'" json:"passWord"`
	HeadPortrait string `gorm:"type:varchar(35);comment:'头像名称'" json:"headPortrait"`
	NickName     string `gorm:"type:varchar(20);comment:'昵称';unique_index" json:"nickName"`
	Email        string `gorm:"type:varchar(30);comment:'邮箱'" json:"email"`
	Tel          string `gorm:"type:varchar(11);comment:'电话'" json:"tel"`
	QuestionA    string `gorm:"type:varchar(30);comment:'密保问题A'" json:"questionA"`
	QuestionB    string `gorm:"type:varchar(30);comment:'密保问题B'" json:"questionB"`
	QuestionC    string `gorm:"type:varchar(30);comment:'密保问题C'" json:"questionC"`
	Flag         int    `gorm:"type:int(0);default:-1;comment:'用户类型标志位'" json:"flag"`
}

// TableName 实现 Tabler 接口来更改默认表名，例如：
//func (User) TableName() string {
//	return "profiles"
//}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("Find After")
	return
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}
	return
}
