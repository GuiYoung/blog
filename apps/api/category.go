package api

import (
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCate create category
func CreateCate(c *gin.Context) {
	var cate models.Category
	_ = c.ShouldBindJSON(&cate)
	var code int
	if code = models.CheckCate(cate.Name); code == errmsg.ERROR_CATE_NOT_EXIST {
		code = models.CreateCate(&cate)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateInfo get cate info
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	cate, code := models.GetOneCate(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateList get cate list
func GetCateList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	cateList, code, count := models.GetCateList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"cateList": cateList,
		"count":    count,
		"message":  errmsg.GetErrMsg(code),
	})
}

// EditCateInfo edit cate info
func EditCateInfo(c *gin.Context) {
	var cate models.Category
	_ = c.ShouldBindJSON(&cate)

	code := models.EditCateInfo(&cate)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}

// DestroyCate destroy cate
func DestroyCate(c *gin.Context) {
	var cate models.Category
	_ = c.ShouldBindJSON(&cate)

	code := models.DestroyOneCate(&cate)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}
