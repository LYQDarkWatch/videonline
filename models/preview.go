package models

type Preview struct {
	Video_ID   int    ` json:"video_id"`
	Video_Name string `json:"video_name"`
	TagID      int    ` json:"tag_id" gorm:"primary_key"`
	//Tag           Tag    `json:"video_tag" gorm:"ForeignKey:Tag_Id"`
	Tag           Tag    `gorm:"foreignkey:TagID"`
	Video_Content string `json:"video_content"`
	Video_Imgurl  string `json:"video_imgurl"`
	Play_Sum      int    `json:"play_sum"`
	Star_Sum      int    `json:"star_sum"`
	Video_Isvip   int    `json:"video_isvip"`
}

//var  priview Preview

//获取所有免费电影
func GetAllFreePreview() (freeVideo []Preview) {
	db.Where("video_isvip = ?", "0").Find(&freeVideo)
	return
}

//获取所有收费电影
func GetAllVipPreview() (vipVideo []Preview) {
	db.Where("video_isvip = ?", "1").Find(&vipVideo)
	return
}

//按分类检索视频
func GetVideoByTag(id int) (previewbytag []Preview) {
	db.Preload("Tag").Where("tag_id = ?", id).Find(&previewbytag)
	return previewbytag
}

//搜索视频
func SearchVideoByName(name string) (search []Preview) {
	db.Where("video_name like ?", "%"+name+"%").Find(&search)
	return
}

//新增视频预览
func AddPreview(video_name, video_content, video_img string, tag_id int) bool {
	db.Create(&Preview{
		Video_Name:    video_name,
		TagID:         tag_id,
		Video_Content: video_content,
		Video_Imgurl:  video_img,
		Play_Sum:      0,
		Star_Sum:      0,
		Video_Isvip:   0,
	})
	return true
}

//免费电影变VIP电影
func FreeVideoBeVIP(video_id int) {
	db.Where("video_id = ?", video_id).Model(&Preview{}).UpdateColumn("video_isvip", 1).First(&preview)
}

//VIP视频变为免费视频
func VIPVideoBeFree(video_id int) {
	db.Where("video_id = ?", video_id).Model(&Preview{}).UpdateColumn("video_isvip", 0).First(&preview)
}
