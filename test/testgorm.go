package main

import (
	"ginchat/http/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/ginchat?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&model.UserBasic{})

	// Create
	user := &model.UserBasic{}
	user.Name = "jz.wang"
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	user.HeartbeatTime = time.Now()
	db.Create(user)
	db.Model(user).Update("PassWord", "1234")
	// Read
	//   db.First(&product, 1) // 根据整型主键查找
	//   db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	//   // Update - 将 product 的 price 更新为 200
	//   db.Model(&product).Update("Price", 200)
	//   // Update - 更新多个字段
	//   db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//   db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//   // Delete - 删除 product
	//   db.Delete(&product, 1)
}
