package api

import (
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddPost add post
func AddPost(c *gin.Context) {
	var post models.Post
	_ = c.ShouldBindJSON(&post)
	var code int

	code = models.AddPost(&post)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    post,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeletePost delete post
func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := models.DeletePost(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetPostList get post list
func GetPostList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	postList, code, count := models.GetPostList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"cateList": postList,
		"count":    count,
		"message":  errmsg.GetErrMsg(code),
	})
}

// GetPost get post
func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, code := models.GetOnePost(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"cate":    post,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditPostInfo edit post
func EditPostInfo(c *gin.Context) {
	var post models.Post
	_ = c.ShouldBindJSON(&post)

	code := models.EditPost(&post)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetPostsCate get posts of one cate
func GetPostsCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	posts, code := models.GetPostsCate(uint(id), pageSize, pageNum)

	count := len(posts)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"posts":   posts,
		"count":   count,
		"message": errmsg.GetErrMsg(code),
	})
}
