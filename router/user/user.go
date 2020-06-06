package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/login/controller"
)

// Init 初始化路由
func Init(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/register", controller.UserCreate)
		user.GET("/:id", controller.UserDetail)
	}
}
