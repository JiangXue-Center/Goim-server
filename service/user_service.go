package service

import (
	"Goim-server/models"
	"Goim-server/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var secretKey = []byte("hf-goim-server")

func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	fmt.Println(data)
	c.JSON(200, gin.H{
		"message": data,
	})
}

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user == nil {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if !utils.ValidPassword(password, user.Salt, user.Password) {
		c.JSON(404, gin.H{
			"message": "账号或密码错误",
		})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID

	// fmt.Println("密钥:", utils.Config.GetString("secretKey"))
	tokenStr, err := token.SignedString([]byte(utils.Config.GetString("secretKey")))
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{
		"token": tokenStr,
	})
}

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")

	if data := models.FindUserByName(user.Name); data != nil {
		c.JSON(400, gin.H{
			"message": "用户名已注册",
		})
		return
	}

	password := c.PostForm("password")
	rePassword := c.PostForm("repassword")
	if password != rePassword {
		c.JSON(400, gin.H{
			"message": "两次输入的密码不一致",
		})
		return
	}
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		fmt.Println("Failed to generate salt: %v", err)
	}
	user.Salt = salt
	user.Password = utils.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("id:", id)
	user.ID = uint(id)
	fmt.Println("userId:", user.ID)

	models.DeleteUser(uint(user.ID))
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "修改失败",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "修改参数不匹配",
		})
		return
	}

	user.ID = uint(id)
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改成功",
	})
}

func FindUser(c *gin.Context) {
	name := c.Query("name")
	user := models.FindUserByName(name)
	c.JSON(200, gin.H{
		"data": user,
	})
}
