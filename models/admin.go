package models

type Admin struct {
	Admin_ID int `gorm:"primary_key" json:"id"`
	Admin_Name string `json:"admin_name"`
	Admin_Passwd string `json:"admin_passwd"`
	Created_Time string `json:"create_time"`
	Priority int `json:"priority"`
}

var admin Admin
func CheckAdmin(username, password string) bool {
	db.Select("admin_id").Where(Admin{Admin_Name:username,Admin_Passwd:password}).First(&admin)
	if admin.Admin_ID > 0{
		return true
	}
	return false
}

func CreateAdmin(username,password,createdtime string) bool {
	if ExistAdminByName(username) == false{
		db.Create(&Admin{
			Admin_Name:username,
			Admin_Passwd:password,
			Created_Time:createdtime,
		})
		return true
	}
	return false
}
func ExistAdminByName(username string) bool {
	if result := db.Select("admin_name").Where("admin_name = ?", username).First(&admin).Error
	result != nil {
		return false
	}
	return true
}

func BecomeVip(username string) bool {
	if result := db.Where("admin_name = ?",username).Model(&Admin{}).UpdateColumn("priority",1).First(&admin).Error
	result != nil{
		return false
	}
	return true
}