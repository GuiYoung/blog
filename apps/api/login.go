package api

import (
	"blog/models"
	"blog/utils/errmsg"
	"blog/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login login
func Login(c *gin.Context) {
	// get username and pwd
	var user models.User
	_ = c.ShouldBindJSON(&user)

	// check login
	data, code := models.CheckLogin(&user)

	token := ""
	// set token
	token, err := jwt.GenerateToken(data.UserName)
	if err != nil {
		code = errmsg.ERROR
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"ID":       data.ID,
		"userName": data.UserName,
		"message":  errmsg.GetErrMsg(code),
		"token":    token,
	})
}

// FrontLogin front login
func FrontLogin(c *gin.Context) {
	// get username and pwd
	var user models.User
	_ = c.ShouldBindJSON(&user)

	// check login
	data, code := models.CheckFrontLogin(&user)

	token := ""
	// set token
	token, err := jwt.GenerateToken(data.UserName)
	if err != nil {
		code = errmsg.ERROR
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"ID":       data.ID,
		"userName": data.UserName,
		"message":  errmsg.GetErrMsg(code),
		"token":    token,
	})
}
