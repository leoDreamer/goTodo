package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 定义配置文件路径
type Config struct {
	Name string
}

// InitConfig 获取配置文件
func (c *Config) InitConfig() error {
	if c.Name == "" {
		panic(fmt.Errorf("config name required"))
	}
	viper.SetConfigFile(c.Name)
	viper.SetConfigType("json")

	return viper.ReadInConfig()
}
