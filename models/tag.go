package models

type Tag struct {
	Tag_Id       int
	Tag_Name     string
	Created_Time string
}

var tag Tag

func GetTags() (tags []Tag) {
	db.Find(&tags)
	return
}

//func GetTagTotal(maps interface{}) (count int) {
//	db.Model(&Tag{}).Where(maps).Count(&count)
//	return
//}

//添加标签
func AddTag(name string, Created_time string) bool {
	db.Create(&Tag{
		Tag_Name:     name,
		Created_Time: Created_time,
	})
	return true
}

//根据id查找标签
func FindTagBYID(tag_name string) int {
	db.Model(&tag).Where("tag_name = ?", tag_name).First(&tag)
	id := tag.Tag_Id
	return id
}

//验证标签名称是否已经存在
func ExistTagByName(name string) bool {
	if result := db.Select("tag_id").Where("tag_name = ?", name).First(&tag).Error; result != nil {
		return false
	}
	return true
}

//验证次id标签是否存在
func ExistTagByID(id string) bool {
	var tag Tag
	if result := db.Select("tag_id").Where("tag_id=?", id).First(&tag).Error; result != nil {
		return false
	}
	return true
}

//删除标签
func DeleteTag(id string) bool {
	db.Where("tag_id = ?", id).Delete(&Tag{})
	return true
}

//func EditTag(id int,data interface{}) bool {
//	db.Model(&Tag{}).Where("id= ?",id).Update(data)
//	return true
//}

//这属于 gorm 的 Callbacks ，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除
//时将被调用，如果任何回调返回错误，gorm将停止未来操作并回滚所有更改。
//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
//	 scope.SetColumn("CreatedOn", time.Now().Unix())
//	 return nil
//	 }
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	 scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//	}
