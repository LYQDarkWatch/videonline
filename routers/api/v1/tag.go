package v1

import (
	"time"

	//"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"videOnline/models"
	"videOnline/pkg/error"
)

//获取多个文章标签
func GetTags(c *gin.Context) {

	data := make(map[string]interface{})
	data["lists"] = models.GetTags()
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	var tags models.Tag
	c.BindJSON(&tags)

	name := tags.Tag_Name
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	createdtime := time.Format("2006-1-02 15:04:05")
	code := error.INVALID_PARAMS
	if models.ExistTagByName(name) == false {
		models.AddTag(name, createdtime)
		code = error.SUCCESS
	} else {
		code = error.ERROR_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

////修改文章标签
//func EditTag(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//	name := c.Query("name")
//	modifiedBy := c.Query("modified_by")
//
//	valid := validation.Validation{}
//
//	var state int = -1
//	if arg := c.Query("state"); arg != ""{
//		state = com.StrTo(arg).MustInt()
//		valid.Range(state,0,1,"state").Message("状态只允许0或1")
//	}
//	valid.Required(id, "id").Message("ID不能为空")
//	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
//	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
//	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
//
//	code := error.INVALID_PARAMS
//
//	if ! valid.HasErrors(){
//		code = error.SUCCESS
//		if models.ExistTagByID(id){
//			data := make(map[string]interface{})
//			data["modified_by"] = modifiedBy
//			if name != ""{
//				data["name"] = name
//			}
//			if state != -1 {
//				data["state"] = state
//			}
//			models.EditTag(id,data)
//		} else {
//			for _, err := range valid.Errors{
//				log.Println(err.Key,err.Message)
//			}
//		}
//		c.JSON(http.StatusOK,gin.H{
//			"code":code,
//			"msg":error.GetMsg(code),
//			"data":make(map[string]string),
//		})
//	}
//}
//
//删除文章标签
func DeleteTag(c *gin.Context) {
	id := c.Query("tag_id")
	code := error.INVALID_PARAMS
	if models.ExistTagByID(id) == true {
		models.DeleteTag(id)
		code = error.SUCCESS
	} else {
		code = error.ERROR_NOT_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}
