package handlers

import (
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

func Login(c *gin.Context) {
	var json struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User == "admin" && json.Password == "admin" {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Hello World",
	})
}
