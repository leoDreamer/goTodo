package middleware

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		n, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("body params:", string(n))
		c.Next()
	}
}
