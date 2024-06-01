package main

import (
	"Goim-server/app/user/internal/database"
	"Goim-server/app/user/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitMySqlDB()
	// 创建一个默认的 Gin 路由器
	router := gin.Default()
	routes.SetupRoutes(router)
	// 启动并运行服务器
	router.Run(":8080")
}
