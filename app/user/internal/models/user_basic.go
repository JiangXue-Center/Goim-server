package models

import (
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogoutTime    *time.Time
	IsLogout      bool
	DeviceInfo    string
}
