package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gitlab.com/login/config"
	"gitlab.com/login/middleware"
	"gitlab.com/login/router"
)

func main() {
	// 记录日志
	f, _ := os.Create(viper.GetString("log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 初始化路由
	r := router.InitRuter()
	r.Use(middleware.Exception())
	r.Use(middleware.Logger())
	r.Run()
}
