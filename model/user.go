package model

import db "gitlab.com/login/dbs"

// User 用户model
type User struct {
	ID     int64  `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Passwd string `json:"password" form:"password"`
}

// Create 数据库插入用户
func (u *User) Create() (err error) {
	exists, errE := db.Conns.Exec("SELECT COUNT(name) FROM user WHERE name=?", u.Name)
	if errE != nil {
		panic(err)
	}
	res, err := db.Conns.Exec("INSERT INTO user(name, password) VALUES (?, ?)", u.Name, u.Passwd)
	if err != nil {
		panic(err)
	}
	u.ID, err = res.LastInsertId()

	return err
}
