package server

import (
	"Goim-server/app/user/internal/models"
	"Goim-server/app/user/internal/models/request"
	"Goim-server/common/cache"
	"Goim-server/common/database"
	"Goim-server/common/response"
	"Goim-server/common/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

func RegisterUser(req *request.RegisterRequest) response.Response {
	// 检查两次输入的密码是否一致
	if req.Password != req.ConfirmPassword {
		return response.Response{
			Message: "password does not match",
			Code:    http.StatusBadRequest,
			Data:    "",
		}
	}
	//生成哈希密码
	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		log.Println("Error hashing password", err)
		return response.Response{
			Message: "Internal server error",
			Code:    http.StatusInternalServerError,
			Data:    "",
		}
	}
	// 将用户信息存储到数据库
	user := models.User{
		Name:     req.Username,
		Password: hashedPassword,
	}
	db := database.MysqlDB
	if err := db.Create(&user).Error; err != nil {
		return response.Response{
			Message: "Error creating user",
			Code:    http.StatusInternalServerError,
			Data:    "",
		}
	}
	return response.Response{
		Message: "User registered successfully",
		Code:    http.StatusOK,
		Data:    "",
	}
}

func LoginUser(req *request.LoginRequest) response.Response {
	var user models.User
	if err := database.MysqlDB.Where("name = ?", req.Username).First(&user).Error; err != nil {
		return response.Response{
			Message: "User not found",
			Code:    http.StatusNotFound,
			Data:    "",
		}
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return response.Response{
			Message: "Password does not match",
			Code:    http.StatusBadRequest,
			Data:    "",
		}
	}

	token, err := utils.GenerateJWT(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		return response.Response{
			Message: "Failed to generate token",
			Code:    http.StatusInternalServerError,
			Data:    "",
		}
	}
	err = cache.SetValue(cache.RedisVal.LoginUserId(user.ID), token, 7*24*time.Hour)
	if err != nil {
		return response.Response{
			Message: "Internal server error",
			Code:    http.StatusInternalServerError,
			Data:    "",
		}
	}
	return response.Response{
		Message: "User logged in",
		Code:    http.StatusOK,
		Data:    token,
	}
}
