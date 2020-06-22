package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/login/middleware"
	"gitlab.com/login/router/user"
)

// InitRuter 初始化路由
func InitRuter() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(middleware.Exception())
	route.Use(middleware.Logger())
	user.Init(route)

	return route
}
