package dbs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Conns mysql连接实例
var Conns *sql.DB

func init() {
	var err error
	// username := viper.GetString("username")
	// password := viper.GetString("password")
	port := viper.GetString("port")
	db := viper.GetString("db")
	host := viper.GetString("host")

	dns := fmt.Sprintf("tcp(%s:%s)/%s?parseTime=true", host, port, db)

	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = Conns.Ping()
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
