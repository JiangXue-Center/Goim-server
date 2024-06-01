package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email    string `json:"email" valid:"email"`
	Avatar   string `json:"avatar"`
	Salt     string `json:"salt"`
}

// TableName 指定表名为 `user`
func (User) TableName() string {
	return "user"
}

func main() {
	// 设置配置文件名和路径
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config") // 指定配置文件的目录

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 打印读取到的配置
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))

	// 从配置文件中读取数据库连接信息
	dsn := viper.GetString("mysql.dsn")
	if dsn == "" {
		log.Fatalf("Database DSN not found in config")
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
