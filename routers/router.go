package routers

import (
	"github.com/gin-gonic/gin"
	"videOnline/middleware/cors"
	"videOnline/middleware/jwt"
	"videOnline/pkg/setting"
	"videOnline/routers/api"
	v1 "videOnline/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	//默认加载了Recovery()的中间件，所以在不知道如何处理error的时候，可以直接panic出去
	r.Use(gin.Recovery())
	r.Use(cors.Cors())

	gin.SetMode(setting.RunMode)
	r.POST("/login", api.GetAdmin)
	r.POST("/register", api.CreateAdmin)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/admin/editadmin", api.EditAdminInfo)
		//升级会员
		apiv1.POST("/bevip", api.BecomeVip)
		//获取所有标签
		apiv1.GET("/tags/getalltags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags/addtag", v1.AddTag)
		////更新指定标签
		//apiv1.PUT("/tags/:id", v1.EditTag)
		////删除指定标签
		apiv1.DELETE("/tags/deletetag", v1.DeleteTag)
		//获取视频列表
		apiv1.GET("videos", v1.GetVideos)
		////获取指定视频
		//apiv1.GET("videos/:id",v1.GetVideo)
		//新建视频
		apiv1.POST("/videos", v1.AddVideo)
		//更新视频信息
		//apiv1.PUT("videos/:id",v1.EditTag)
		//删除视频
		apiv1.DELETE("/videos/:id", v1.DeleteVideo)
		//获取文章列表
		//apiv1.GET("/articles",v1.GetArticles)
		////获取指定文章
		//apiv1.GET("/articles/:id",v1.GetArticle)
		////新建文章
		//apiv1.POST("/articles",v1.AddArticle)
		////更新指定文章
		//apiv1.PUT("articles/:id",v1.EditArticle)
		////删除指定文章
		//apiv1.DELETE("/articles/:id",v1.DeleteArticle)

		//播放指定视频
		apiv1.GET("/video/play", v1.StreamHandler)
		apiv1.POST("/video/upload/:vid-id", v1.UploadHandler)
		//获取所有免费视频
		apiv1.GET("/video/getallfree", v1.GetAllFreePriview)
		//获取所有VIP视频
		apiv1.GET("/video/getallvip", v1.GetAllVipPriview)
		//获取单个视频详情
		apiv1.GET("/video/getvideo", v1.GetVideoByID)
		//根据分类检索视频
		apiv1.GET("video/getvideobytag", v1.GetVideosByTag)
	}
	return r
}
