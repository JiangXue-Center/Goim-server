package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个默认的 Gin 路由器
    router := gin.Default()

    // 定义一个带参数的 GET 路由
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(http.StatusOK, gin.H{
            "user": name,
        })
    })

    // 定义一个带查询参数的 GET 路由
    router.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname")

        c.JSON(http.StatusOK, gin.H{
            "firstname": firstname,
            "lastname":  lastname,
        })
    })

    // 定义一个 POST 路由
    router.POST("/login", func(c *gin.Context) {
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
    })

    // 启动并运行服务器
    router.Run(":8080")
}