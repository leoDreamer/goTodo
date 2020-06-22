package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/login/errors"
	"gitlab.com/login/model"
	"gitlab.com/login/service"
)

// UserCreate 创建用户
func UserCreate(c *gin.Context) {
	var err error

	m := new(model.User)
	if err = c.ShouldBindJSON(&m); err != nil {
		errors.ErrFn(errors.PARAMERR, err)
	}
	err = m.Create()
	if err != nil {
		panic(err)
	}

	service.R(c, service.StructToMap(m))

}

// UserDetail 创建用户
func UserDetail(c *gin.Context) {}

//
