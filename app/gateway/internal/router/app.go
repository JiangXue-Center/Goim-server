package router

import (
	"Goim-server/app/gateway/internal/handler"
	"Goim-server/service"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/list", service.GetUserList)
	r.POST("/login", service.Login)
	r.POST("/user", service.CreateUser)
	r.DELETE("/user/:id", service.DeleteUser)
	r.PUT("/user/:id", service.UpdateUser)

	authGroup := r.Group("/")
	authGroup.Use(handler.AuthMiddleware())

	authGroup.GET("/user", service.FindUser)
	return r
}
