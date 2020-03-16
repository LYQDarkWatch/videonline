package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"videOnline/models"
	"videOnline/pkg/error"
	"videOnline/pkg/util"
)

type admin struct {
	Admin_ID          string
	Admin_Name        string `valid:"Required; MaxSize(50)"`
	Admin_Passwd      string `valid:"Required; MaxSize(50)"`
	Admin_display     string
	Admin_CreatedTime string
	Priority          int
}

type notification struct {
	Admin_name string `json:"admin_name"`
	User_name  string `json:"user_name"`
	User_id    string `json:"user_id"`
}

var b admin
var notifi notification

//管理员登录
func GetAdmin(c *gin.Context) {
	c.BindJSON(&b)
	username := b.Admin_Name
	password := b.Admin_Passwd
	println("username:", username)
	println("password:", password)
	data := make(map[string]interface{})
	code := error.INVALID_PARAMS
	if models.CheckAdmin(username, password) == true {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = error.ERROR_AUTH_TOKEN
		} else {
			data["user"] = models.GetAdminInfo(username)
			data["token"] = token
			code = error.SUCCESS
		}
	} else {
		code = error.ERROR_USER_LOGIN
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//管理员获取所有用户
func AdminGetAllUser(c *gin.Context) {
	data := make(map[string]interface{})
	data["user"] = models.GetAllUser()
	var code int
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//禁止用户评论
func DisableUserComments(c *gin.Context) {
	c.BindJSON(&notifi)
	admin_name := notifi.Admin_name
	println("admin name:", admin_name)
	id := notifi.User_id
	println("user_id :", id)
	user_id, _ := strconv.Atoi(id)
	user_name := notifi.User_name
	println("user_name :", user_name)
	content := "您的账号因发表违禁言论已经被限制评论，七天后解封"
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	sendtime := time.Format("2006-1-02 15:04:05")
	var code int
	println(user_id)
	if models.DisableUserComments(admin_name, user_name, content, sendtime, user_id) == true {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//恢复用户评论功能
func AbleUserComments(c *gin.Context) {
	c.BindJSON(&notifi)
	id := notifi.User_id
	admin_name := notifi.Admin_name
	user_id, _ := strconv.Atoi(id)
	user_name := notifi.User_name
	content := "您的账号已被解封"
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	sendtime := time.Format("2006-1-02 15:04:05")
	var code int
	println(user_id)

	if models.AbleUserComments(admin_name, user_name, content, sendtime, user_id) == true {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//管理员获取所有视频信息
func AdminGetAllVideoInfo(c *gin.Context) {
	data := make(map[string]interface{})
	data["videos"], data["previews"] = models.AdminGetAllVideo()
	var code int
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//管理员查看现有标签
func AdminGetAllTag(c *gin.Context) {
	var code int
	data := make(map[string]interface{})
	data["tags"] = models.GetTags()
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//管理员删除视频
func AdminDeleteVideo(c *gin.Context) {
	var code int
	vid := c.Query("video_id")
	video_id, _ := strconv.Atoi(vid)
	if models.AdminDeleteVideo(video_id) == true {
		code = error.SUCCESS
	} else {
		code = error.ERROR_VIDEO_DELETE_ERROE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

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
