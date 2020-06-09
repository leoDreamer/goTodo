package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	fmt.Printf("start config\n")
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
