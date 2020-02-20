package models

type Preview struct {
	Video_ID      int    `gorm:"primary_key" json:"video_id"`
	Video_Name    string `json:"video_name"`
	Tag_ID        int    `json:"tag_id"`
	Priview_Tag   Tag    `json:"video_tag" gorm:"foreignkey:tag_id"`
	Video_Content string `json:"video_content"`
	Video_Imgurl  string `json:"video_imgurl"`
	Video_ImgUrl  string `json:"video_img_url"`
	Video_HotSum  int    `json:"video_hot_sum"`
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
func GetVideoByTag(id string) (priviewbytag []Preview) {
	db.Where("tag_id = ?", id).Preload("Priview_Tag").Find(&priviewbytag)
	return
}
