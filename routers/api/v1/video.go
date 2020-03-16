package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"videOnline/models"
	"videOnline/pkg/error"
	"videOnline/pkg/setting"
)

//获取单个视频详情
func GetVideoByID(c *gin.Context) {

	id := c.Query("vid")
	println(id)
	code := error.INVALID_PARAMS
	var data interface{}
	//if models.ExistVideoByID(id) == true {
	models.VideoPlaySum(id)
	data = models.GetVideoByID(id)
	code = error.SUCCESS
	//} else {
	//	code = error.ERROR_NOT_EXIST_ARTICLE
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//按分类获取多个视频
func GetVideosByTag(c *gin.Context) {
	id := c.Query("tag_id")
	k, _ := strconv.Atoi(id)
	data := make(map[string]interface{})
	data["lists"] = models.GetVideoByTag(k)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//获取多个视频
func GetVideos(c *gin.Context) {

}

//获取所有免费视频
func GetAllFreePriview(c *gin.Context) {
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := error.SUCCESS
	data["lists"] = models.GetAllFreePreview()
	//data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//获取所有VIP视频
func GetAllVipPriview(c *gin.Context) {
	data := make(map[string]interface{})
	code := error.SUCCESS
	data["lists"] = models.GetAllVipPreview()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

type videos struct {
	Video_Name    string `json:"video_name"`
	Video_Content string `json:"video_content"`
	Video_Info    string `json:"video_info"`
	Tag_Name      string `json:"tag_name"`
	Video_Actor   string `json:"video_actor"`
	Video_Url     string `json:"video_url"`
	Video_Imgurl  string `json:"video_imgurl"`
}

//新增视频
func AddVideo(c *gin.Context) {

	var video videos
	c.BindJSON(&video)

	name := video.Video_Name
	content := video.Video_Content
	info := video.Video_Info
	tag_name := video.Tag_Name
	actor := video.Video_Actor
	video_img := video.Video_Imgurl
	video_playurl := video.Video_Url
	println("play:", video_playurl)
	println("tag_name", tag_name)
	tag_id := models.FindTagBYID(tag_name)
	println("tag_id: ", tag_id)
	timeNow := time.Now().Unix()
	time := time.Unix(timeNow, 0)
	createdtime := time.Format("2006-1-02 15:04:05")

	models.AddVideo(name, info, video_playurl, actor, createdtime, tag_id)
	models.AddPreview(name, content, video_img, tag_id)
	var code int
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改视频信息
func EditVideo(c *gin.Context) {

}

//管理员删除视频
func AdminDeleteVideo(c *gin.Context) {
	var code int
	id := c.Query("video_id")
	video_id, _ := strconv.Atoi(id)
	models.AdminDeleteVideo(video_id)
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

const (
	VIDEO_DIR = "D:/videos/"
)

//搜索视频
func SearchVideo(c *gin.Context) {
	name := c.Query("video_name")
	data := make(map[string]interface{})
	//code := error.INVALID_PARAMS
	models.SearchVideoByName(name)
	//if models.ExistVideoByName(name) == true{
	data["list"] = models.SearchVideoByName(name)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//播放视频
func StreamHandler(c *gin.Context) {
	available := setting.TokenBucket.TakeAvailable(1)
	if available <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "超出播放限制",
		})

	} else {
		vname := c.Query("vname")
		vl := VIDEO_DIR + vname
		println(vl)
		video, err := os.Open(vl)

		if err != nil {
			log.Println(err)
			return
		}
		videos, _ := ioutil.ReadAll(video)
		c.Data(200, "video/mp4", videos)
		defer video.Close()
	}
}

//免费视频变VIP视频
func FreeVideoBeVip(c *gin.Context) {
	id := c.Query("video_id")
	video_id, _ := strconv.Atoi(id)
	var code int
	models.FreeVideoBeVIP(video_id)
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//VIP视频变免费视频
func VipVideoBeFree(c *gin.Context) {
	id := c.Query("video_id")
	video_id, _ := strconv.Atoi(id)
	var code int
	models.VIPVideoBeFree(video_id)
	code = error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
	})
}

//获取当前热门视频
func GetHotVideo(c *gin.Context) {
	data := make(map[string]interface{})
	code := error.SUCCESS
	data["lists"] = models.GetHotPreview()
	//data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//获最新视频
func GetNewVideo(c *gin.Context) {
	data := make(map[string]interface{})
	code := error.SUCCESS
	data["news"] = models.GetNewPreview()
	//data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}
