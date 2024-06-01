package handlers

import (
	"Goim-server/app/user/internal/models/request"
	accountservice "Goim-server/app/user/internal/server/accountservice"
	"Goim-server/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"user": name,
	})
}

func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSONResponse(c, "fail to format JSON", http.StatusBadRequest, "")
		return
	}
	resp := accountservice.RegisterUser(&req)
	response.JSONResponse(c, resp.Message, resp.Code, resp.Data)
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSONResponse(c, "fail to format JSON", http.StatusBadRequest, "")
	}
	resp := accountservice.LoginUser(&req)
	response.JSONResponse(c, resp.Message, resp.Code, resp.Data)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Hello World",
	})
}
