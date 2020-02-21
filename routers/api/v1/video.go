package v1

import (
	//"github.com/astaxie/beego/validation"
	//"fmt"
	//"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	//"github.com/juju/ratelimit"

	//"github.com/juju/ratelimit"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"time"

	//"videOnline/middleware/bucketoken"
	"videOnline/models"
	"videOnline/pkg/error"
	"videOnline/pkg/setting"
)

//获取单个视频详情
func GetVideoByID(c *gin.Context) {

	id := c.Query("vid")
	println(id)
	//valid := validation.Validation{}
	//valid.Min(id, 1, "vid").Message("ID必须大于0")

	code := error.INVALID_PARAMS
	var data interface{}
	//if ! valid.HasErrors() {
	//	if models.ExistVideoByID(id) == true{
	//		data = models.GetVideoByID(id)
	//		code = error.SUCCESS
	//	} else {
	//		code = error.ERROR_NOT_EXIST_ARTICLE
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		fmt.Println(err.Key, err.Message)
	//	}
	//}
	if models.ExistVideoByID(id) == true {
		data = models.GetVideoByID(id)
		code = error.SUCCESS
	} else {
		code = error.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//按分类获取多个视频
func GetVideosByTag(c *gin.Context) {
	id := c.Query("tag_id")
	data := make(map[string]interface{})
	data["lists"] = models.GetVideoByTag(id)
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
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := error.SUCCESS
	data["lists"] = models.GetAllVipPreview()
	//data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

//新增视频
func AddVideo(c *gin.Context) {

}

//修改视频信息
func EditVideo(c *gin.Context) {

}

//删除视频
func DeleteVideo(c *gin.Context) {

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
	//bucket := bucketoken.NewConnLimiter(2)
	//if bucket.GetToken(1) == false{
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 500,
	//		"msg":  "超出播放限制",
	//	})
	//}

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
		//c.Header("Content-Type","video/mp4")
		c.Data(200, "video/mp4", videos)
		//bucket.ReleaseToken()
		defer video.Close()
	}

}

//func StreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
//	vid := p.ByName("video-id")
//	vl := VIDEO_DIR + vid
//	video, err := os.Open(vl)
//	if err != nil{
//		print("error")
//		return
//	}
//	w.Header().Set("Content-Type","video/mp4")
//	http.ServeContent(w,r,"",time.Now(), video)
//	defer video.Close()
//}
func UploadHandler(c *gin.Context) {

}

//func IScunzai(c *gin.Context) {
//	id := c.Query("vid")
//	println(id)
//	if models.ExistVideoByID(id) == true {
//		c.JSON(http.StatusOK, gin.H{
//			"code": "200",
//			"msg":  "此视频存在",
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": "400",
//			"msg":  "此视频不存在",
//		})
//	}
//}
