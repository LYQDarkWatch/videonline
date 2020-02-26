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
	if models.ExistFavorite(u_id, v_id) == false {
		if models.AddFavorite(u_id, v_id, video_name, add_time) == true {
			models.AddVideoStarSum(video_id)
			code = error.SUCCESS
		}
	} else {
		code = error.ERROR_FAVORITE_EXIST
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

//删除收藏的视频
func DeleteFavoriteByID(c *gin.Context) {
	uid := c.Query("user_id")
	vid := c.Query("video_id")
	u_id, _ := strconv.Atoi(uid)
	v_id, _ := strconv.Atoi(vid)
	code := error.INVALID_PARAMS
	if models.DeleteFavoriteByID(u_id, v_id) == true {
		code = error.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}
