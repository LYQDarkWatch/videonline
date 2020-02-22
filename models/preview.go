package models

type Preview struct {
	Video_ID   int    ` json:"video_id"`
	Video_Name string `json:"video_name"`
	TagID      int    ` json:"tag_id" gorm:"primary_key"`
	//Tag           Tag    `json:"video_tag" gorm:"ForeignKey:Tag_Id"`
	Tag           Tag    `gorm:"foreignkey:TagID"`
	Video_Content string `json:"video_content"`
	Video_ImgUrl  string `json:"video_img_url"`
	Play_Sum      int    `json:"play_sum"`
	Star_Sum      int    `json:"star_sum"`
	Video_IsVip   int    `json:"video_is_vip"`
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
	//db.Model(&previewbytag).AddForeignKey("tag_id","tag(tag_id)","RESTRICT","RESTRICT")
	//db.Where("tag_id = ?", id).Preload("Tag").Find(&previewbytag)
	db.Preload("Tag").Where("tag_id = ?", id).Find(&previewbytag)

	//db.Model(&previewbytag).Preload("Priview_Tag").Where("tag_id=?",id).Find(&previewbytag)
	//db.Preload("Priview_Tag",id).Find(&previewbytag).Where("tag_id=?",id)
	return previewbytag
}

//搜索视频
func SearchVideoByName(name string) (search []Preview) {
	db.Where("video_name like ?", "%"+name+"%").Find(&search)
	return
}

func GetTagInPreview() (previewbytag []Preview) {
	db.Preload("Tag").Find(&previewbytag)
	return
}
