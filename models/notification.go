package models

type Notification struct {
	Notification_Id int    `json:"notification_id"`
	Admin_name      string `json:"admin_name"`
	User_id         int    `json:"user_id"`
	User_name       string `json:"user_name"`
	Content         string `json:"content"`
	Send_Time       string `json:"send_time"`
}

var notification Notification

//封禁用户评论功能
func DisableUserComments(admin_name, user_name, content, send_time string, user_id int) bool {
	db.Where("user_id = ?", user_id).Model(&User{}).UpdateColumn("can_comment", 1).First(&user)
	db.Create(&Notification{
		Admin_name: admin_name,
		User_id:    user_id,
		User_name:  user_name,
		Content:    content,
		Send_Time:  send_time,
	})
	return true
}

//恢复用户评论功能
func AbleUserComments(admin_name, user_name, content, send_time string, user_id int) bool {
	db.Where("user_id = ?", user_id).Model(&User{}).UpdateColumn("can_comment", 0).First(&user)
	db.Create(&Notification{
		Admin_name: admin_name,
		User_id:    user_id,
		User_name:  user_name,
		Content:    content,
		Send_Time:  send_time,
	})
	return true
}

//发送通知给用户
func SentNotiToUser(admin_name, user_name, content, send_time string, user_id int) bool {
	db.Create(&Notification{
		Admin_name: admin_name,
		User_id:    user_id,
		User_name:  user_name,
		Content:    content,
		Send_Time:  send_time,
	})
	return true
}

//获取用户的通知
func UserGetNoti(username string) (notifis []Notification) {
	println("user:", username)
	db.Where("user_name = ?", username).Order("notification_id desc").Find(&notifis)
	return
}

//删除通知
func DeleteUserGetNoti(nofi_id int) bool {
	println("user:", nofi_id)
	db.Where("notification_id = ?", nofi_id).Delete(Notification{})
	return true
}
