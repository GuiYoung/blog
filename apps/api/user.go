package api

import (
	"blog/models"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)
	var code int

	msg, err := validator.Validate(&user)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	if code = models.CheckUserExit(&user); code == errmsg.ERROR_USER_NOT_EXIST {
		code = models.CreateUser(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePwd change password
func ChangePwd(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)

	code := models.ChangePwd(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// ChangeUserInfo change info
func ChangeUserInfo(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)

	code := models.ChangeUserInfo(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserInfo get user
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := models.GetUser(uint(id))

	var user = make(map[string]interface{})
	user["userName"] = data.UserName
	user["authCode"] = data.AuthCode

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"user":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserList get users
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	userList, code, count := models.GetUserList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"userList": userList,
		"count":    count,
		"message":  errmsg.GetErrMsg(code),
	})
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)

	code := models.DeleteUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
