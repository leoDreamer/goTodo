package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println("body params: \n %s", string(n))
		c.Next()
	}
}
