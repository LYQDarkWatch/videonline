package v1

import (
	"time"

	//"github.com/Unknwon/com"
	"net/http"
	"videOnline/models"
	"videOnline/pkg/error"

	"github.com/gin-gonic/gin"
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

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := c.Query("tag_id")
	code := error.INVALID_PARAMS
	if models.ExistTagByID(id) == true {
		if models.DeleteTag(id) == true {
			code = error.SUCCESS
		} else {
			code = error.ERROR_TAG_IS_QUOTE
		}

	} else {
		code = error.ERROR_NOT_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}
