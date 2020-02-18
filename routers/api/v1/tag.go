package v1

import (
	"log"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"videOnline/models"
	"videOnline/pkg/error"
	"videOnline/pkg/setting"
	"videOnline/pkg/util"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := error.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 10, "name").Message("名称最长10字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长允许100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = error.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = error.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			}
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg"  : error.GetMsg(code),
		"data" : make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := error.INVALID_PARAMS

	if ! valid.HasErrors(){
		code = error.SUCCESS
		if models.ExistTagByID(id){
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != ""{
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id,data)
		} else {
			for _, err := range valid.Errors{
				log.Println(err.Key,err.Message)
			}
		}
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":error.GetMsg(code),
			"data":make(map[string]string),
		})
	}
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id").Message("ID必须大于0")
	code := error.INVALID_PARAMS
	if ! valid.HasErrors(){
		code = error.SUCCESS
		if models.ExistTagByID(id){
			models.DeleteTag(id)
		}else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	}else {
		for _,err := range valid.Errors{
			log.Println(err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg": error.GetMsg(code),
		"data": make(map[string]string),
	})
}
