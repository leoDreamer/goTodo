package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/login/model"
)

// UserCreate 创建用户
func UserCreate(c *gin.Context) {
	var err error
	m := new(model.User)
	if err = c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = m.Create()
	if err != nil {
		panic(err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   m,
	})

}

// UserDetail 创建用户
func UserDetail(c *gin.Context) {}

//
