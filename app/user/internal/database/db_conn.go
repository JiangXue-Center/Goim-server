package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDB *gorm.DB

func InitMySqlDB() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config") // 指定配置文件的目录

	var err error

	// 读取配置文件
	err = viper.ReadInConfig()
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
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}
