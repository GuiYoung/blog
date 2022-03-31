package models

import (
	"blog/databases"
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type PostTag struct {
	gorm.Model
	PostID uint `gorm:"type:bigint" json:"postID"`
	TagID  uint `gorm:"type:bigint" json:"tagID"`
}

// AddPostTag add
func AddPostTag(tagID, postID uint) (code int) {
	if err := databases.db.Create(PostTag{TagID: tagID, PostID: postID}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeletePostTag delete
func DeletePostTag(postTag *PostTag) (code int) {
	if err := databases.db.Delete(&postTag).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// FindPostsByTag find many by tagID
func FindPostsByTag(tagID uint) (posts []Post, code int) {
	if err := databases.db.Where("id = ?", tagID).Find(&posts).Error; err != nil {
		return posts, errmsg.ERROR
	}
	return posts, errmsg.SUCCESS
}

// FindTagByPost find many by postID
func FindTagByPost(postID uint) (tags []Tag, code int) {
	if err := databases.db.Where("id = ?", postID).Find(&tags).Error; err != nil {
		return tags, errmsg.ERROR
	}
	return tags, errmsg.SUCCESS
}

// update
