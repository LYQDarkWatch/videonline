package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"videOnline/models"
	"videOnline/pkg/error"
)

type contents struct {
	Video_ID      string
	User_ID       string
	User_Display  string
	User_Logo     string
	Video_Content string
}

var content contents

//添加评论
func AddContent(c *gin.Context) {
	c.BindJSON(&content)
	video_id := content.Video_ID
	user_id := content.User_ID
	user_name := content.User_Display
	user_logo := content.User_Logo
	video_content := content.Video_Content
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	add_time := time.Format("2006-1-02 15:04:05")
	v_id, _ := strconv.Atoi(video_id)
	u_id, _ := strconv.Atoi(user_id)
	code := error.INVALID_PARAMS
	if models.AddContent(v_id, u_id, user_name, user_logo, video_content, add_time) == true {
		code = error.SUCCESS
	} else {
		code = error.ERROR_ADD_CONTENT_ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//删除评论
func DeleteContent(c *gin.Context) {
	cid := c.Query("content_id")
	uid := c.Query("user_id")
	//vid := c.Query("video_id")
	content_id, _ := strconv.Atoi(cid)
	user_id, _ := strconv.Atoi(uid)
	//video_id,_ := strconv.Atoi(vid)

	code := error.INVALID_PARAMS
	if models.DeleteContent(content_id, user_id) == true {
		code = error.SUCCESS
	} else {
		code = error.ERROR_DELETE_CONTENT_NOT_MYSELF
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}
