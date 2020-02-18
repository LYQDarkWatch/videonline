package models

type Video struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
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
//func ExistArticleByID(id int) bool  {
//	var article Article
//	db.Select("id").Where("id=?").First(&article)
//	if article.ID >0 {
//		return true
//	}
//	return false
//}
//
//func GetArticleTotal(maps interface{}) (count int) {
//	db.Model(&Article{}).Where(maps).Count(&count)
//	return
//}

