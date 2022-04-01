package models

import (
	"blog/utils/errmsg"
	"errors"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100)" json:"name"`
	Desc  string `gorm:"type:varchar(200)" json:"desc"`
	Count int    `gorm:"type:int" json:"count"`
}

// CheckTag check tag
func CheckTag(name string) (code int) {
	tag := Tag{}
	if err := Db.Where("name = ?", name).First(&tag).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.ERROR_TAG_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// AddTag create tag
func AddTag(tag *Tag) (code int) {
	if Db.Where("name = ?", tag.Name).First(&Tag{}).RowsAffected > 0 {
		return errmsg.ERROR_TAG_USED
	}
	if err := Db.Create(tag).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteTag delete tag
func DeleteTag(tag *Tag) (code int) {
	if err := Db.Delete(&tag).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetOneTag get tag
func GetOneTag(id uint) (tag Tag, code int) {
	if err := Db.Where("id = ?", id).First(&tag).Error; err != nil {
		return tag, errmsg.ERROR
	}
	return tag, errmsg.SUCCESS
}

// GetTags get tags
func GetTags() (tags []Tag, code int, count int) {
	if err := Db.Find(&tags).Error; err != nil {
		return tags, errmsg.ERROR, count
	}
	count = len(tags)
	return tags, errmsg.SUCCESS, count
}

// EditTag edit tag
func EditTag(tag *Tag) (code int) {
	if err := Db.Model(&Category{}).Where("id = ?", tag.ID).Updates(Tag{Name: tag.Name, Desc: tag.Desc}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UpdateCount update count
func (tag Tag) UpdateCount() (int, error) {
	var posts []Post
	if err := Db.Where("id = ?", tag.ID).Find(&posts).Error; err != nil {
		return len(posts), err
	}
	return len(posts), nil
}
