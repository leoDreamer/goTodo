package model

import (
	"fmt"

	db "gitlab.com/login/dbs"
)

// User 用户model
type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

// Create 数据库插入用户
func (u *User) Create() (err error) {
	existsCount := 0
	errE := db.Conns.QueryRow("SELECT COUNT(name) FROM user WHERE name=?", u.Name).Scan(&existsCount)
	fmt.Printf("%d", existsCount)
	if errE != nil {
		panic(err)
	}
	if existsCount > 0 {
		panic("exits user")
	}
	res, err := db.Conns.Exec("INSERT INTO user(name, password) VALUES (?, ?)", u.Name, u.Password)
	if err != nil {
		panic(err)
	}
	u.ID, err = res.LastInsertId()

	return err
}
