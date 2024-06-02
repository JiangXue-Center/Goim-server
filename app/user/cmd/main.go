package main

import (
	"Goim-server/app/user/routes"
	"Goim-server/common/cache"
	"Goim-server/common/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitMySqlDB()
	cache.InitRedisConn()
	// 创建一个默认的 Gin 路由器
	router := gin.Default()
	routes.SetupRoutes(router)
	// 启动并运行服务器
	router.Run(":8080")
}
