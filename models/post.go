package models

import (
	"blog/databases"
	"blog/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Post struct {
	Category Category `gorm:"foreignkey:CateID"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Summary string `gorm:"type:varchar(255);not null" json:"summary"`
	Content string `gorm:"type:longtext" json:"content"`
	ViewNum int    `gorm:"type:int" json:"viewNum"`
	CateID  uint   `gorm:"type:bigint" json:"cateID"`
}

// AddPost add post
func AddPost(post *Post) (code int) {
	if err := databases.db.Create(post).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeletePost delete post
func DeletePost(id int) (code int) {
	if err := databases.db.Delete(&Post{}, id).Error; err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetOnePost get post
func GetOnePost(id uint) (post Post, code int) {
	if err := databases.db.Where("id = ?", id).First(&post).Error; err != nil {
		return post, errmsg.ERROR
	}
	databases.db.Model(&Post{}).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return post, errmsg.SUCCESS
}

// GetPostList get postList
func GetPostList(pageSize, pageNum int) (posts []Post, code int, count int) {
	if err := databases.db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&posts).Error; err != nil {
		return posts, errmsg.ERROR, count
	}
	count = len(posts)
	return posts, errmsg.SUCCESS, count
}

// EditPost edit post
func EditPost(post *Post) (code int) {

	if err := databases.db.Model(&Post{}).Where("id = ?", post.ID).Updates(Post{Title: post.Title, Summary: post.Summary, Content: post.Content, CateID: post.CateID}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetPostsCate get posts of one cate
func GetPostsCate(id uint, pageSize, pageNum int) (posts []Post, code int) {
	cate := Category{}
	if err := databases.db.Where("id = ?", id).First(&cate).Error; err != nil {
		return posts, errmsg.ERROR_CATE_NOT_EXIST
	}

	if err := databases.db.Where("cate_id = ?", id).Find(&posts).Error; err != nil {
		return posts, errmsg.ERROR
	}

	return posts, errmsg.SUCCESS
}

// get posts of one tag

// search post by title
//func SearchPostsByTitle(title string, pageSize, pageNum int) (posts []Post, code int) {
//
//}
