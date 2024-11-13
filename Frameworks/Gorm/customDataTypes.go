package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"slices"
	"strings"
)

// Gorm自定义数据在数据库和 Go 结构体之间转换的行为
// 对于自定义类型的转换，需要实现两个接口：Scanner 和 Valuer
// Scanner 接口: 定义从数据库中读取数据时的自定义行为。
// Valuer 接口: 定义将 Go 结构体中的数据写入数据库时的自定义行为。
// 注意实现的Scan和Value方法需要对数据库字段类型进行转换，
// 具体的转换类型需根据具体需求进行选择。

// Grade 自定义类型
// 表示军官等级
type Grade struct {
	General    bool
	Colonel    bool
	Captain    bool
	Lieutenant bool
}

// Scan 实现 Scanner 接口
// 完成从数据库读取数据时的自定义行为
func (g *Grade) Scan(src interface{}) error {
	valueRaw, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	valueStr := string(valueRaw)
	valueStr = strings.Trim(valueStr, "{}")
	values := strings.Split(valueStr, ", ")
	if slices.Contains(values, "General") {
		g.General = true
	}

	if slices.Contains(values, "Colonel") {
		g.Colonel = true
	}

	if slices.Contains(values, "Captain") {
		g.Captain = true
	}

	if slices.Contains(values, "Lieutenant") {
		g.Lieutenant = true
	}
	return nil
}

// Value 实现 Valuer 接口
// 完成将 Go 结构体中的数据写入数据库时的自定义行为
func (g *Grade) Value() (driver.Value, error) {
	values := make([]string, 0, 4)
	if g.General {
		values = append(values, "General")
	}
	if g.Colonel {
		values = append(values, "Colonel")
	}
	if g.Captain {
		values = append(values, "Captain")
	}
	if g.Lieutenant {
		values = append(values, "Lieutenant")
	}
	return "{" + strings.Join(values, ", ") + "}", nil
}

// User 结构体
type User struct {
	ID     uint64 `gorm:"primaryKey"`
	Name   string
	Grades *Grade `gorm:"type:text"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Learn-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic(err)
	//}

	//user1 := User{
	//	Name: "yz",
	//	Grades: &Grade{
	//		General:    true,
	//		Colonel:    true,
	//		Captain:    true,
	//		Lieutenant: true,
	//	},
	//	ID: 1,
	//}
	//
	//user2 := User{
	//	Name: "yz2",
	//	Grades: &Grade{
	//		General:    true,
	//		Colonel:    true,
	//		Captain:    false,
	//		Lieutenant: false,
	//	},
	//	ID: 2,
	//}
	//
	//// 插入多个数据
	//err = db.Create(&[]User{user1, user2}).Error
	//if err != nil {
	//	panic(err)
	//}

	// 读取数据转化为Json
	var users []User
	err = db.Find(&users).Error
	if err != nil {
		panic(err)
	}

	jsonBytes, err := MarshalJSON(users)
	if err != nil {
		panic(err)
	}
	println(string(jsonBytes))
}

func MarshalJSON(users []User) ([]byte, error) {
	buf := &strings.Builder{}
	encoder := json.NewEncoder(buf)
	for _, user := range users {
		err := encoder.Encode(user)
		if err != nil {
			return nil, err
		}
	}
	return []byte(buf.String()), nil
}
