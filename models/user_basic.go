package models

import (
	"Goim-server/utils"
	"fmt"
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

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) *UserBasic {
	user := UserBasic{}
	if err := utils.DB.Where("name = ?", name).First(&user).Error; err != nil {
		// 处理查询过程中的错误
		return nil
	}
	return &user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func CreateUser(user UserBasic) *gorm.DB {
	// utils.DB.Find()

	return utils.DB.Create(&user)
}

func DeleteUser(userId uint) *gorm.DB {
	return utils.DB.Delete(&UserBasic{}, userId)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Email: user.Email, Phone: user.Phone})
}
