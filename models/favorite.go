package models

type Favorite struct {
	Favorite_Id int `json:"favorite_id"`
	User_ID     int `json:"user_id"`
	//User_Name     string  `json:"user_name"`
	Video_ID      int     `json:"video_id"`
	Video_Name    string  `json:"video_name"`
	Tag_ID        int     `json:"tag_id"`
	Favorite_Time string  `json:"favorite_time"`
	Preview       Preview `gorm:"ForeignKey:Video_ID;AssociationForeignKey:Video_ID"`
}

var favorite Favorite

//将视频添加到我的收藏
func AddFavorite(user_id, video_id int, video_name, favorite_time string) bool {
	db.Create(&Favorite{
		User_ID:       user_id,
		Video_ID:      video_id,
		Video_Name:    video_name,
		Favorite_Time: favorite_time,
	})
	return true
}

//检查是否已经存在该收藏
func ExistFavorite(user_id, video_id int) bool {
	if result := db.Select("favorite_id").Where("user_id = ? and video_id = ?", user_id, video_id).First(&favorite).Error; result != nil {
		return false
	}
	return true
}

//获取收藏夹视频
func GetUserFavorite(id int) (allfavorite []Favorite) {
	db.Where("user_id=?", id).Preload("Preview").Preload("Preview.Tag").Find(&allfavorite)
	return allfavorite
}

//删除收藏夹视频
func DeleteFavoriteByID(favorite_id int) bool {
	result := db.Where("favorite_id = ?", favorite_id).Delete(&favorite).Error
	if result != nil {
		return false
	}
	return true
}
