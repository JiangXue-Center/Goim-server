package models

import (
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
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
