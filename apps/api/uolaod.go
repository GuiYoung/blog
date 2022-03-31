package api

import (
	"blog/service"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	code := errmsg.SUCCESS

	url, err := service.UpLoadFile(file, fileSize)
	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
