package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.SetConfigName("app")
	Config.AddConfigPath("config")
	err := Config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", Config.Get("app"))
	fmt.Println("config mysql", Config.Get("mysql"))

}

func InitMySQL() {
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(Config.GetString("mysql.dns")), &gorm.Config{Logger: log})
}
