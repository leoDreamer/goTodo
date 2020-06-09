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
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	port := viper.GetString("mysql.port")
	db := viper.GetString("mysql.db")
	host := viper.GetString("mysql.host")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, db)

	fmt.Println("%s", dns)

	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = Conns.Ping()
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
