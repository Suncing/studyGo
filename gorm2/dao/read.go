package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm2/gorm2/db"
	"gorm2/gorm2/model"
	"gorm2/gorm2/model/itemvideo"
	"gorm2/gorm2/utils"
	"log"
)

// GetOne 获取第一条记录（主键升序）
func GetOne() {
	user := itemvideo.User{}
	db.DB.First(&user)
	utils.ToString(user)
}

// GetOne2 获取一条记录，没有指定排序字段
func GetOne2() {
	user := itemvideo.User{}
	db.DB.Take(&user)
	utils.ToString(user)
}

// GetLast 获取最后一条记录（主键降序
func GetLast() {
	user := itemvideo.User{}
	db.DB.Last(&user)
	utils.ToString(user)
}

// GetNotFoundErr 检查 ErrRecordNotFound 错误
func GetNotFoundErr() {
	user := itemvideo.User{}
	err := db.DB.Where("nick_name=?", "lucy").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("没有找到该条记录！")
	}
	utils.ToString(user)
}

// GetAll 获取全部
func GetAll() {
	user := make([]itemvideo.User, 0)
	result := db.DB.Find(&user)
	log.Fatalln(result.RowsAffected)
}

// GetA 获取一个字段
func GetA() {
	//user := make([]string, 0)
	var name string
	db.DB.Model(itemvideo.User{}).Where("id=?", 1).Select("user_name").Find(&name)
	fmt.Println(name)

	var name1 string
	db.DB.Model(itemvideo.User{}).Where("id=?", 1).Pluck("user_name", &name1)
	fmt.Println(name1)

	var names []string
	db.DB.Model(itemvideo.User{}).Where("user_name != ?", "NULL").Select("user_name").Find(&names)
	fmt.Println(names)

	var names2 []string
	db.DB.Model(itemvideo.User{}).Where("user_name != ?", "NULL").Pluck("user_name", &names2)
	fmt.Println(names2)
}

//条件查询
/**
// 获取第一条匹配的记录
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
// 获取全部匹配的记录
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';
// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');
// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';
// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
*/

// GetStruct Struct条件
func GetStruct() {
	user := make([]itemvideo.User, 0)
	db.DB.Where(&itemvideo.User{NickName: "lisa"}).Find(&user)
	utils.ToString(user)
}

// GetMap GetMap条件
func GetMap() {
	user := make([]itemvideo.User, 0)
	db.DB.Where(map[string]interface{}{"nick_name": "lisa"}).Find(&user)
	utils.ToString(user)
}

// GetSlice 主键切片查询
func GetSlice() {
	user := make([]itemvideo.User, 0)
	db.DB.Where([]uint{2, 3, 4}).Find(&user)
	utils.ToString(user)
}

func GetOffset() {
	users := make([]itemvideo.User, 0)
	db.DB.Limit(10).Offset(3).Find(&users)
	utils.ToString(users)
}

//注意 当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件，
//可以使用 map 来构建查询条件，例如：
//db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)

// GetGroup Group & Having
func GetGroup() {
	users := make([]itemvideo.User, 0)
	db.DB.Select("avg(pass_word)").Group("user_name").Find(&users)
	utils.ToString(users)
}

// GetDistinct  去重
func GetDistinct() {
	var userName []string
	db.DB.Model(itemvideo.User{}).Where("user_name != ?", "NULL").Distinct("user_name").Pluck("user_name", &userName)
	fmt.Println(userName)
}

// GetJoins Join 普通Find
func GetJoins() {
	temp := struct {
		model.Student
		model.Achievement
	}{}
	db.DB.Model(model.Student{}).
		Select("students.*,achievements.*").
		Joins("INNER JOIN achievements on students.id = achievements.student_id").
		Find(&temp)
	marshal, _ := json.Marshal(temp)
	fmt.Println(string(marshal))

}

// GetJoinsScan Join Scan形式
func GetJoinsScan() {
	temp := struct {
		model.Student
		model.Achievement
	}{}
	db.DB.Model(model.Student{}).
		Select("students.*,achievements.*").
		Joins("INNER JOIN achievements on students.id = achievements.student_id").
		Scan(&temp)
	marshal2, _ := json.Marshal(temp)
	fmt.Println(string(marshal2))
}

// GetJoinsRows Join Rows形式
func GetJoinsRows() {
	rows, _ := db.DB.Model(model.Student{}).
		Select("students.*,achievements.*").
		Joins("INNER JOIN achievements on students.id = achievements.student_id").
		Rows()
	for rows.Next() {
		a := struct {
			model.Student
			model.Achievement
		}{}
		db.DB.ScanRows(rows, &a)
		fmt.Println(a)
	}
}
