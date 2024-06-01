package response

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(c *gin.Context, message string, code int, data interface{}) {
	c.JSON(code, Response{
		Message: message,
		Code:    code,
		Data:    data,
	})
}
