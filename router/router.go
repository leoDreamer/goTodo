package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/login/router/user"
)

// InitRuter 初始化路由
func InitRuter() *gin.Engine {
	route := gin.Default()
	user.Init(route)

	return route
}
