package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)
//获取单个视频
func GetVideo(c *gin.Context)  {

}
//获取多个视频
func GetVideos(c *gin.Context)  {

}
//新增视频
func AddVideo(c *gin.Context)  {

}
//修改视频信息
func EditVideo(c *gin.Context)  {

}
//删除视频
func DeleteVideo(c *gin.Context)  {

}

const  (
	VIDEO_DIR = "./videos/"
)
func StreamHandler(c *gin.Context){
	vid := c.Query("vid")
	vl := VIDEO_DIR + vid
	print(vl)
	video, err := os.Open("./videos/"+vid)

	if err != nil{
		print("error")
		return
	}
	videos, _ := ioutil.ReadAll(video)
	//c.Header("Content-Type","video/mp4")
	c.Data(200,"video/mp4",videos)
	defer video.Close()
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
func UploadHandler(c *gin.Context){

}