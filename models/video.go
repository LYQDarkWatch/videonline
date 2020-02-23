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
	Video_Logo   string    `json:"video_logo"`
	First_Upload string    `json:"first_upload"`
	Last_Update  string    `json:"last_update"`
	Play_Sum     int       `json:"play_sum"`
	Star_Sum     int       `json:"star_sum"`
	Content_Sum  int       `json:"content_sum"`
	Content      []Content `gorm:"foreignkey:video_id"`
}

//func (article *Article) BeforeCreate(scope gorm.Scope) error {
//	scope.SetColumn("CreateOn",time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope gorm.Scope) error {
//	scope.SetColumn("ModifiedOn",time.Now().Unix())
//
//	return nil
//}
//
var video Info
var preview Preview

//获取单个视频详情
func GetVideoByID(id string) (video Info) {
	db.Where("video_id = ?", id).Preload("Video_Tag").Preload("Content").First(&video)
	return
}

//访问视频增加访问总数
func VideoPlaySum(id string) bool {
	//db.Select("video_id").Where("video_id = ?", id).First(&video).Exec("UPDATE video_info SET play_sum = play_sum + 1")
	//db.Select("video_id").Where("video_id = ?", id).First(&preview).Exec("UPDATE video_preview SET play_sum = play_sum + 1")
	db.Model(&video).Where("video_id = ?", id).Update("play_sum", gorm.Expr("play_sum + 1"))
	db.Model(&preview).Where("video_id = ?", id).Update("play_sum", gorm.Expr("play_sum + 1"))
	return true
}

//收藏数增加
func VideoStarSum(id string) bool {
	db.Select("video_id").Where("video_id = ?", id).First(&video).Exec("UPDATE video_info SET star_sum = star_sum + 1")
	db.Select("video_id").Where("video_id = ?", id).First(&preview).Exec("UPDATE video_preview SET star_sum = star_sum + 1")
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

//func GetArticleTotal(maps interface{}) (count int) {
//	db.Model(&Article{}).Where(maps).Count(&count)
//	return
//}
