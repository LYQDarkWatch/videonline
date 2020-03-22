package models

import "github.com/jinzhu/gorm"

type Info struct {
	Video_Id     int       `gorm:"primary_key" json:"video_id"`
	Video_Name   string    `json:"video_name"`
	Video_Info   string    `json:"video_info"`
	Tag_ID       int       `json:"tag_id"`
	Video_Tag    Tag       `gorm:"ForeignKey:Tag_ID;AssociationForeignKey:Tag_ID"`
	Video_Url    string    `json:"video_url"`
	Video_Actor  string    `json:"video_actor"`
	First_Upload string    `json:"first_upload"`
	Last_Update  string    `json:"last_update"`
	Play_Sum     int       `json:"play_sum"`
	Star_Sum     int       `json:"star_sum"`
	Content_Sum  int       `json:"content_sum"`
	Content      []Content `gorm:"foreignkey:video_id"`
	Video_Isvip  int       `json:"video_isvip"`
}

var video Info
var preview Preview

//获取单个视频详情
func GetVideoByID(id string) (video Info) {
	db.Where("video_id = ?", id).Preload("Video_Tag").Preload("Content").First(&video)
	return
}

//管理员查看所有视频
func AdminGetAllVideo() (video []Info, preview []Preview) {
	//db.Model(&video).Preload("Video_Tag").Preload("Content").Find(&video)
	db.Model(&preview).Preload("Tag").Preload("Info").Preload("Info.Video_Tag").Preload("Info.Content").Find(&preview)
	return
}

//访问视频增加访问总数
func VideoPlaySum(id string) bool {
	db.Model(&video).Where("video_id = ?", id).Update("play_sum", gorm.Expr("play_sum + 1"))
	db.Model(&preview).Where("video_id = ?", id).Update("play_sum", gorm.Expr("play_sum + 1"))
	return true
}

//收藏数增加
func AddVideoStarSum(id string) bool {
	db.Model(&video).Where("video_id = ?", id).Update("star_sum", gorm.Expr("star_sum + 1"))
	db.Model(&preview).Where("video_id = ?", id).Update("star_sum", gorm.Expr("star_sum + 1"))
	return true
}

//检查视频是否存在
func ExistVideoByID(id string) bool {
	if result := db.Select("video_id").Where("video_id=?", id).First(&video).Error; result != nil {
		return false
	}
	return true
}

//检查视频是否存在
func ExistVideoByName(name string) bool {
	if result := db.Select("video_name").Where("video_name=?", name).First(&video).Error; result != nil {
		return false
	}
	return true
}

//新建视频
func AddVideo(video_name, video_info, video_url, actor, created_time string, tag_id int) bool {
	println("video_url:", video_url)
	db.Create(&Info{
		Video_Name:   video_name,
		Video_Info:   video_info,
		Tag_ID:       tag_id,
		Video_Url:    video_url,
		Video_Actor:  actor,
		First_Upload: created_time,
		Last_Update:  "0",
		Play_Sum:     0,
		Star_Sum:     0,
		Content_Sum:  0,
		Content:      nil,
	})
	return true
}

//管理员删除视频
func AdminDeleteVideo(video_id int) bool {
	error := db.Where("video_id = ?", video_id).Delete(&Info{}).Error
	if error != nil {
		return false
	} else {
		error1 := db.Where("video_id = ?", video_id).Delete(&Preview{}).Error
		if error1 != nil {
			return false
		} else {
			return true
		}
	}
}

//func GetArticleTotal(maps interface{}) (count int) {
//	db.Model(&Article{}).Where(maps).Count(&count)
//	return
//}
