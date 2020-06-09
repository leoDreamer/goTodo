package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exception 错误拦截中间件
func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
			}
		}()
		c.Next()
	}
}
