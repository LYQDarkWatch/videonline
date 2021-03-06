package models

type Admin struct {
	Admin_ID     int    `gorm:"primary_key"`
	Admin_Name   string `json:"admin_name"`
	Admin_Passwd string `json:"admin_passwd"`
	Priority     int    `json:"priority"`
}

var admin Admin

func BanUsercomment(user_id int, user_name string) {
	db.Where("user_name = ? and user_name = ?", user_id, user_name).Model(&User{}).UpdateColumn("can_comment", 1).First(&user)
}

//管理员登录
func CheckAdmin(admin_name, admin_pass string) bool {

	result := db.Where("admin_name = ? AND admin_passwd = ?", admin_name, admin_pass).First(&Admin{}).Error
	if result != nil {
		return false
	}
	return true
}

//获取管理员详情
func GetAdminInfo(admin_name string) (admin Admin) {
	db.Where("admin_name = ?", admin_name).First(&admin)
	return
}
