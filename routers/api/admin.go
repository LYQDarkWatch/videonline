package api

//
//import (
//	"github.com/astaxie/beego/validation"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"time"
//	"videOnline/models"
//	"videOnline/pkg/error"
//	"videOnline/pkg/util"
//)
//
//type admin struct {
//	Admin_ID          string
//	Admin_Name        string `valid:"Required; MaxSize(50)"`
//	Admin_Passwd      string `valid:"Required; MaxSize(50)"`
//	Admin_display     string
//	Admin_CreatedTime string
//	Priority          int
//}
//
//var a admin
//
////注册新用户
//func CreateAdmin(c *gin.Context) {
//	c.BindJSON(&a)
//	username := a.Admin_Name
//	password := a.Admin_Passwd
//	displayname := a.Admin_display
//	timeNow := time.Now().Unix()
//	time := time.Unix(timeNow, 0)
//	createdtime := time.Format("2006-1-02 15:04:05")
//	//valid := validation.Validation{}
//	a = admin{Admin_Name: username, Admin_Passwd: password, Admin_CreatedTime: createdtime}
//	//ok,_ := valid.Valid(&a)
//	data := make(map[string]interface{})
//	code := error.INVALID_PARAMS
//	//if ok {
//	isCreate := models.CreateAdmin(username, password, displayname, createdtime)
//	if isCreate == true {
//		data["token"] = "created success"
//		code = error.SUCCESS
//	}
//	//}else {
//	//	for _,err := range valid.Errors{
//	//		log.Println(err.Key,err.Message)
//	//	}
//	//}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  error.GetMsg(code),
//		"data": data,
//	})
//}
//
////用户登录
//func GetAdmin(c *gin.Context) {
//	c.BindJSON(&a)
//	username := a.Admin_Name
//	password := a.Admin_Passwd
//	valid := validation.Validation{}
//	a = admin{Admin_Name: username, Admin_Passwd: password}
//	ok, _ := valid.Valid(&a)
//
//	data := make(map[string]interface{})
//	code := error.INVALID_PARAMS
//
//	if ok {
//		isExist := models.CheckAdmin(username, password)
//		if isExist {
//			token, err := util.GenerateToken(username, password)
//			if err != nil {
//				code = error.ERROR_AUTH_TOKEN
//			} else {
//				data["token"] = token
//				code = error.SUCCESS
//			}
//		} else {
//			code = error.ERROR_AUTH
//		}
//	} else {
//		for _, err := range valid.Errors {
//			log.Println(err.Key, err.Message)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  error.GetMsg(code),
//		"data": data,
//	})
//}
//
////修改用户资料
//func EditAdminInfo(c *gin.Context) {
//	c.BindJSON(&a)
//	id := a.Admin_ID
//	password := a.Admin_Passwd
//	displayname := a.Admin_display
//	data := make(map[string]interface{})
//	data["admin_display"] = displayname
//	data["admin_passwd"] = password
//	code := error.INVALID_PARAMS
//	if models.ExistAdminByID(id) == true {
//		if models.EditAdminInfo(displayname, id, data) == true {
//			code = error.SUCCESS
//		} else {
//			code = error.ERROR_NOT_SAME_ADMIN
//		}
//	} else {
//		code = error.ERROR_ADMIN_NOT_EXIST
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  error.GetMsg(code),
//	})
//}
//
////升级成为会员
//func BecomeVip(c *gin.Context) {
//	c.BindJSON(&a)
//	username := a.Admin_Name
//	code := error.INVALID_PARAMS
//	println(username)
//	if models.BecomeVip(username) == true {
//		code = error.SUCCESS
//		c.JSON(http.StatusOK, gin.H{
//			"code": code,
//			"msg":  error.GetMsg(code),
//			"data": "恭喜成为我们的VIP用户，感谢您的支持",
//		})
//	} else {
//		code = error.INVALID_PARAMS
//		c.JSON(http.StatusOK, gin.H{
//			"code": code,
//			"msg":  error.GetMsg(code),
//			"data": "升级会员失败，请稍后再试",
//		})
//	}
//
//}
