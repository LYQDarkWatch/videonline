package models

type Favorite struct {
	User_ID       int     `json:"user_id"`
	Video_ID      int     `gorm:"primary_key" json:"video_id"`
	Video_Name    string  `json:"video_name"`
	Tag_ID        int     `json:"tag_id"`
	Favorite_Time string  `json:"favorite_time"`
	Preview       Preview `json:"Preview" gorm:"foreignkey:video_id"`
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

//获取收藏夹视频
func GetUserFavorite(id int) (allfavorite []Favorite) {
	db.Where("user_id=?", id).Preload("Preview").Find(&allfavorite)
	return allfavorite
}
