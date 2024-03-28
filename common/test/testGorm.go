package main

import (
	"Goim-server/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql", viper.Get("mysql"))

	db, err := gorm.Open(mysql.Open("root:tencentYun123456@tcp(www.artnecthub.com:3306)/goim-server?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	// user := &models.UserBasic{}
	// user.Name = "Hf"
	// db.Create(user)

	// Read
	// var product Product
	// db.First(&product, 1) // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// fmt.Println(db.First(user, 1))

	// // Update - 将 product 的 price 更新为 200
	// db.Model(user).Update("Password", "123456")
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
