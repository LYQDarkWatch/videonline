package models

import "github.com/jinzhu/gorm"

type Content struct {
	Content_ID    int    `gorm:"primary_key" json:"content_id"`
	Video_ID      int    `json:"video_id"`
	User_ID       int    `json:"user_id"`
	User_Name     string `json:"user_name"`
	User_Logo     string `json:"user_logo"`
	Star_Num      int    `json:"star_num"`
	Video_Content string `json:"video_content"`
	Add_Time      string `json:"add_time"`
}

var content Content
var info Info

//添加评论
func AddContent(vid, uid int, uname, ulogo, vcontent, addtime string) bool {
	db.Model(&info).Where("video_id = ?", vid).Update("content_sum", gorm.Expr("content_sum + 1"))

	db.Create(&Content{
		Video_ID:      vid,
		User_ID:       uid,
		User_Name:     uname,
		User_Logo:     ulogo,
		Star_Num:      0,
		Video_Content: vcontent,
		Add_Time:      addtime,
	})
	return true
}

//删除评论
func DeleteContent(content_id, user_id int) bool {
	if result := db.Model(&content).Select("video_id").Where("content_id=? and user_id=?", content_id, user_id).First(&content).Error; result != nil {
		return false
	}
	db.Model(&info).Where("video_id = ?", content.Video_ID).Update("content_sum", gorm.Expr("content_sum - 1"))
	db.Model(&content).Where("content_id = ? and user_id = ?", content_id, user_id).Delete(&content)
	return true
}
