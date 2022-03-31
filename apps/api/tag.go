package api

import (
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddTag add tag
func AddTag(c *gin.Context) {
	var tag models.Tag
	_ = c.ShouldBindJSON(&tag)
	var code int
	if code = models.CheckTag(tag.Name); code == errmsg.ERROR_CATE_NOT_EXIST {
		code = models.AddTag(&tag)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tag":     tag,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetTagInfo get tag info
func GetTagInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, code := models.GetOneTag(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tag":     tag,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetTags get tags
func GetTags(c *gin.Context) {
	tags, code, count := models.GetTags()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tags":    tags,
		"count":   count,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditTag edit tag
func EditTag(c *gin.Context) {
	var tag models.Tag
	_ = c.ShouldBindJSON(&tag)

	code := models.EditTag(&tag)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tag":     tag,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteTag delete tag
func DeleteTag(c *gin.Context) {
	var tag models.Tag
	_ = c.ShouldBindJSON(&tag)

	code := models.DeleteTag(&tag)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"tag":     tag,
		"message": errmsg.GetErrMsg(code),
	})
}
