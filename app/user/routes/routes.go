package routes

import (
	"Goim-server/app/user/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// HTTP 路由
	router.GET("/user/:name", handlers.GetUser)
	router.GET("/welcome", handlers.Welcome)
	router.POST("/auth/register", handlers.Register)
	router.POST("/auth/login", handlers.Login)
	router.GET("/hello", handlers.Hello)

	// WebSocket 路由
	router.GET("/ws/:id", handlers.WebSocketHandler)
}
