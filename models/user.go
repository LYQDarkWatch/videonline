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
}

var user User

func CheckUser(username, password string) bool {
	db.Select("user_id").Where(User{User_Name: username, User_Passwd: password}).First(&user)
	if user.User_ID > 0 {
		return true
	}
	return false
}

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

func EditUserInfo(id string, data interface{}) bool {
	db.Model(&User{}).Where("user_id= ?", id).Update(data)
	return true
}
func ExistUserByDisplay(displayname string) bool {
	var user1 User
	if result := db.Select("user_display").Where("user_display=?", displayname).First(&user1).Error; result != nil {
		return false
	}
	return true
}
func ExistUserByID(id string) bool {
	if result := db.Select("user_id").Where("user_id=?", id).First(&user).Error; result != nil {
		return false
	}
	return true
}

func ExistUserByName(username string) bool {
	if result := db.Select("user_name").Where("user_name = ?", username).First(&user).Error; result != nil {
		return false
	}
	return true
}

func BecomeVip(username string) bool {
	if result := db.Where("user_name = ?", username).Model(&User{}).UpdateColumn("priority", 1).First(&user).Error; result != nil {
		return false
	}
	return true
}
