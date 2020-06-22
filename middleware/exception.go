package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Exception 错误拦截中间件
func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Print(err, "\n")
				// debug.PrintStack()
				errMsg := fmt.Errorf("%v", err).Error()
				fmt.Print(errMsg, "\n")
				// switch errMsg {
				// case errors.DBERR:
				// 	c.JSON(http.StatusServiceUnavailable, gin.H{"errors": err, "code": 500})
				// case errors.PARAMERR:
				// 	c.JSON(http.StatusBadRequest, gin.H{"errors": err, "code": 400})
				// case errors.FORBIDDEN:
				// 	c.JSON(http.StatusForbidden, gin.H{"errors": err, "code": 403})
				// default:
				// 	c.JSON(http.StatusOK, gin.H{"errors": err, "code": 500})
				// }
			}
		}()
		c.Next()
	}
}
