package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"videOnline/models"
	"videOnline/pkg/error"
)

//添加收藏
func AddFavorite(c *gin.Context) {
	user_id := c.Query("user_id")
	u_id, _ := strconv.Atoi(user_id)
	video_id := c.Query("video_id")
	v_id, _ := strconv.Atoi(video_id)

	video_name := c.Query("video_name")
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	add_time := time.Format("2006-1-02 15:04:05")
	code := error.INVALID_PARAMS
	if models.AddFavorite(u_id, v_id, video_name, add_time) == true {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//查看个人收藏夹
func GetUserFavorite(c *gin.Context) {
	user_id := c.Query("user_id")
	u_id, _ := strconv.Atoi(user_id)
	data := make(map[string]interface{})

	data["list"] = models.GetUserFavorite(u_id)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}
