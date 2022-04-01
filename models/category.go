package models

import (
	"blog/utils/errmsg"
	"errors"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Desc string `gorm:"type:varchar(255);not null" json:"desc"`
}

// CheckCate check cate
func CheckCate(name string) (code int) {
	cate := Category{}
	if err := Db.Where("name = ?", name).First(&cate).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.ERROR_CATE_NOT_EXIST
	}
	return errmsg.ERROR_CATENAME_USED
}

// CreateCate create cate
func CreateCate(cate *Category) (code int) {
	if err := Db.Create(cate).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateList get cate list
func GetCateList(pageSize, pageNum int) (categories []Category, code int, count int) {
	if err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error; err != nil {
		return categories, errmsg.ERROR, count
	}
	count = len(categories)
	return categories, errmsg.SUCCESS, count
}

// GetOneCate get one cate
func GetOneCate(id uint) (cate Category, code int) {
	if err := Db.Where("id = ?", id).First(&cate).Error; err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCESS
}

// EditCateInfo edit cate info
func EditCateInfo(cate *Category) (code int) {
	if err := Db.Model(&Category{}).Where("id = ?", cate.ID).Updates(Category{Name: cate.Name, Desc: cate.Desc}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DestroyOneCate destroy one cate
func DestroyOneCate(cate *Category) (code int) {
	if err := Db.Where("cate_id = ?", cate.ID).Delete(&User{}).Error; err != nil {
		return errmsg.ERROR
	}
	if err := Db.Delete(&cate).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
