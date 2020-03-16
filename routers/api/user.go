package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"videOnline/models"
	"videOnline/pkg/error"
	"videOnline/pkg/util"
)

type user struct {
	User_ID      string
	User_Name    string `valid:"Required; MaxSize(50)"`
	User_Passwd  string `valid:"Required; MaxSize(50)"`
	User_Display string
	Created_Time string
	Priority     int
	User_Phone   string
	User_Email   string
	Login_time   string
}

var a user

//注册新用户
func CreateUser(c *gin.Context) {
	c.BindJSON(&a)
	username := a.User_Name
	password := a.User_Passwd
	displayname := a.User_Display
	phone := a.User_Phone
	email := a.User_Email
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	createdtime := time.Format("2006-1-02 15:04:05")
	code := error.INVALID_PARAMS
	if models.ExistUserByDisplay(displayname) == false {
		if VerifyPhoneFormat(phone) == true {
			if VerifyEmailFormat(email) == true {
				if models.CreateUser(username, password, displayname, createdtime, phone, email) == true {
					code = error.SUCCESS
				} else {
					code = error.ERROR_USER_EXIST
				}
			} else {
				code = error.ERROR_NOT_EMAIL
			}
		} else {
			code = error.ERROR_NOT_PHONE
		}
	} else {
		code = error.ERROT_EXIST_USER_DISPLAY
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//用户登录
func GetUser(c *gin.Context) {
	c.BindJSON(&a)
	username := a.User_Name
	password := a.User_Passwd
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	logintime := time.Format("2006-1-02 15:04:05")
	valid := validation.Validation{}
	a = user{User_Name: username, User_Passwd: password, Login_time: logintime}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := error.INVALID_PARAMS

	if ok {
		isExist := models.CheckUser(username, password, logintime)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = error.ERROR_AUTH_TOKEN
			} else {
				data["user"] = models.GetUserInfo(username)
				data["token"] = token
				code = error.SUCCESS
			}
		} else {
			code = error.ERROR_USER_NOT_EXIST
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//获取用户信息
func GetUserInfo(c *gin.Context) {
	username := c.Query("user_name")
	data := make(map[string]interface{})
	data["user"] = models.GetUserInfo(username)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"info": data,
	})
}

//修改用户资料
func EditUserInfo(c *gin.Context) {
	c.BindJSON(&a)
	id := a.User_ID
	println("id:" + id)
	password := a.User_Passwd
	displayname := a.User_Display
	println(displayname)
	phone := a.User_Phone
	k, _ := strconv.Atoi(phone)
	email := a.User_Email
	data := make(map[string]interface{})
	code := error.INVALID_PARAMS
	if VerifyPhoneFormat(phone) == true {
		if VerifyEmailFormat(email) == true {
			data["user_display"] = displayname
			data["user_passwd"] = password
			data["user_phone"] = k
			data["user_email"] = email
			if models.ExistUserByDisplay(displayname) == false {
				if models.EditUserInfo(id, data) == true {
					code = error.SUCCESS
				}
			} else {
				code = error.ERROR_NOT_SAME_ADMIN
			}
		} else {
			code = error.ERROR_NOT_EMAIL
		}
	} else {
		code = error.ERROR_NOT_PHONE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//升级成为会员
func BecomeVip(c *gin.Context) {
	c.BindJSON(&a)
	username := a.User_Name
	println(username)
	code := error.INVALID_PARAMS
	if models.BecomeVip(username) == true {
		code = error.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  error.GetMsg(code),
			"data": "恭喜成为我们的VIP用户，感谢您的支持",
		})
	} else {
		code = error.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  error.GetMsg(code),
			"data": "升级会员失败，请稍后再试",
		})
	}
}

//e邮箱格式验证
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//手机号验证
func VerifyPhoneFormat(phone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//获取消息提醒
func GetUserNoti(c *gin.Context) {
	user_name := c.Query("user_name")
	data := make(map[string]interface{})
	data["list"] = models.UserGetNoti(user_name)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//获取用户能否评论
func GetUserComment(c *gin.Context) {
	u_id := c.Query("user_id")
	user_id, _ := strconv.Atoi(u_id)
	can_comment := models.GetUserComment(user_id)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": can_comment,
	})
}

//获取用户自己的评论
func UserGetComment(c *gin.Context) {
	u_id := c.Query("user_id")
	user_id, _ := strconv.Atoi(u_id)
	data := make(map[string]interface{})
	data["list"] = models.UserGetContent(user_id)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//删除消息
func DeleteGetUserNoti(c *gin.Context) {
	noti_id := c.Query("notification_id")
	notifi_id, _ := strconv.Atoi(noti_id)
	models.DeleteUserGetNoti(notifi_id)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//删除消息
func GetUserID(c *gin.Context) {
	username := c.Query("user_name")
	user_id := models.GetUserIdByName(username)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"id":   user_id,
	})
}
