package models

import (
	"strconv"
)

type User struct {
	User_ID      int    `gorm:"primary_key" json:"user_id"`
	User_Name    string `json:"user_name"`
	User_Display string `json:"user_display"`
	User_Passwd  string `json:"user_passwd"`
	Created_Time string `json:"create_time"`
	Priority     int    `json:"priority"`
	User_Phone   int    `json:"user_phone"`
	User_Email   string `json:"user_email"`
	Last_Login   string `json:"last_login"`
	Can_Comment  int    `json:"can_comment"`
}

var user User

//获取用户能否评论
func GetUserComment(user_id int) int {
	db.Where("user_id = ?", user_id).First(&user)
	return user.Can_Comment
}

//登录校验
func CheckUser(username, password, time string) bool {
	result := db.Where("user_name = ? AND user_passwd = ?", username, password).First(&User{}).Error
	if result != nil {
		return false
	}
	db.Model(&user).Update("last_login", time)
	return true
}

//注册新用户
func CreateUser(username, password, displayname, createdtime, phone, email string) bool {
	k, _ := strconv.Atoi(phone)
	if ExistUserByName(username) == false {
		db.Create(&User{
			User_Name:    username,
			User_Passwd:  password,
			User_Display: displayname,
			User_Phone:   k,
			User_Email:   email,
			Created_Time: createdtime,
		})
		return true
	}
	return false
}

//更改用户信息
func EditUserInfo(id string, data interface{}) bool {
	db.Model(&User{}).Where("user_id= ?", id).Update(data)
	return true
}

//校验该昵称是否已被占用
func ExistUserByDisplay(displayname string) bool {
	var user1 User
	if result := db.Select("user_display").Where("user_display=?", displayname).First(&user1).Error; result != nil {
		return false
	}
	return true
}

//校验改id用户是否存在
func ExistUserByID(id string) bool {
	if result := db.Select("user_id").Where("user_id=?", id).First(&user).Error; result != nil {
		return false
	}
	return true
}

//校验该用户名是否存在
func ExistUserByName(username string) bool {
	if result := db.Select("user_name").Where("user_name = ?", username).First(&user).Error; result != nil {
		return false
	}
	return true
}

//普通用户升级vip
func BecomeVip(username string) bool {
	if result := db.Where("user_name = ?", username).Model(&User{}).UpdateColumn("priority", 1).First(&user).Error; result != nil {
		return false
	}
	return true
}

//获取用户详情
func GetUserInfo(user_name string) (user User) {
	db.Where("user_name = ?", user_name).First(&user)
	return
}

//管理员获取所有用户
func GetAllUser() (user []User) {
	db.Model(&user).Find(&user)
	return
}

//管理员获取所有用户
func GetUserIdByName(user_name string) int {
	userss := User{}
	db.Select("user_id").Where("user_name=?", user_name).First(&userss)
	println("id", userss.User_ID)
	return userss.User_ID
}
