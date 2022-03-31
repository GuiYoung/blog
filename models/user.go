package models

import (
	"blog/databases"
	"blog/utils/errmsg"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"unsafe"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100)" json:"userName"`
	PassWord string `gorm:"type:varchar(500)" json:"passWord"`
	AuthCode int    `gorm:"type:int" json:"authCode"`
}

// CheckUserExit check user exit
func CheckUserExit(user *User) int {
	data := User{}
	if err := databases.db.Where("id = ?", user.ID).First(&data).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

//// CheckPwd check pwd
//func CheckPwd(user *User) int {
//	data := User{}
//	if err := db.Where("id = ?", user.ID).First(&data).Error; err != nil {
//		return errmsg.ERROR
//	}
//
//	if data.PassWord != Scrypt(user.PassWord) {
//		return errmsg.ERROR_PASSWORD_WRONG
//	}
//
//	return errmsg.SUCCESS
//}

// CreateUser create user
func CreateUser(user *User) int {

	var err error
	if err = databases.db.Where("user_name = ?", user.UserName).First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		user.PassWord = Scrypt(user.PassWord)

		fmt.Println(user.PassWord)
		if err = databases.db.Create(user).Error; err != nil {
			fmt.Println(err)
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

	return errmsg.ERROR_USERNAME_USED

}

// ChangePwd change password
func ChangePwd(user *User) int {
	user.PassWord = Scrypt(user.PassWord)
	fmt.Println(user.PassWord)
	fmt.Println(user.ID)
	if err := databases.db.Model(&User{}).Where("id = ?", user.ID).Update("pass_word", user.PassWord).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ChangeUserInfo change user info
func ChangeUserInfo(user *User) int {
	if err := databases.db.Model(&User{}).Where("id = ?", user.ID).Updates(User{UserName: user.UserName, AuthCode: user.AuthCode}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser find user by id
func GetUser(id uint) (user *User, code int) {
	if err := databases.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// DeleteUser delete user
func DeleteUser(user *User) int {
	if err := databases.db.Delete(&user).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUserList get user list
func GetUserList(pageSize, pageNum int) (users []User, code int, count int) {
	if err := databases.db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error; err != nil {
		return users, errmsg.ERROR, count
	}
	count = len(users)
	return users, errmsg.SUCCESS, count
}

// Scrypt password script
func Scrypt(pwd string) string {
	cost := 10
	dk, _ := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return *(*string)(unsafe.Pointer(&dk))
}

// check login
func CheckLogin(user *User) (*User, int) {
	var data *User

	// check user exits
	if err := databases.db.Where("name = ?", user.UserName).First(&data).Error; err != nil {
		return data, errmsg.ERROR_USER_NOT_EXIST
	}

	// check pwd
	if err := bcrypt.CompareHashAndPassword([]byte(data.PassWord), []byte(user.PassWord)); err != nil {
		return data, errmsg.ERROR_PASSWORD_WRONG
	}

	// check auth code
	if data.AuthCode != 1 {
		return data, errmsg.ERROR_USER_NO_RIGHT
	}

	return data, errmsg.SUCCESS
}

// check fornt login
func CheckFrontLogin(user *User) (*User, int) {
	var data *User

	// check user exits
	if err := databases.db.Where("name = ?", user.UserName).First(&data).Error; err != nil {
		return data, errmsg.ERROR_USER_NOT_EXIST
	}

	// check pwd
	if err := bcrypt.CompareHashAndPassword([]byte(data.PassWord), []byte(user.PassWord)); err != nil {
		return data, errmsg.ERROR_PASSWORD_WRONG
	}

	return data, errmsg.SUCCESS
}
