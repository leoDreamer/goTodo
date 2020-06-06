package main

import (
	"gitlab.com/login/config"
	"gitlab.com/login/middleware"
	"gitlab.com/login/router"
)

func main() {
	// 初始化配置
	config := config.Config{Name: "./config/default.json"}
	config.InitConfig()
	r := router.InitRuter()
	r.Use(middleware.Logger())
	r.Run()
}
