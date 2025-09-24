package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Student 定义结构体用来映射数据库表
type Student struct {
	gorm.Model
	Name string
}

func main() {
	// 建立数据库连接
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移表结构
	db.AutoMigrate(&Student{})

	// 增加数据
	db.Create(&Student{Name: "1"})

	// 查找数据
	var student Student
	db.First(&student, 1)
	db.First(&student, "Name = ?", "1")

	// 更新数据 - update student's price to 200
	db.Model(&student).Update("Name", "1-2")
	// 更新数据 - update multiple fields
	db.Model(&student).Updates(Student{Name: "2"}) // non-zero fields
	db.Model(&student).Updates(map[string]interface{}{"Name": "3"})

	// 删除数据 - delete student
	db.Delete(&student, 1)
}
