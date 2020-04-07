package routers

import (
	"github.com/gin-gonic/gin"
	"videOnline/middleware/cors"
	"videOnline/middleware/jwt"
	"videOnline/middleware/verification"
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
	r.POST("/user/login", api.GetUser)
	r.POST("/user/register", api.CreateUser)
	r.POST("/admin/login", api.GetAdmin)
	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		//修改个人资料
		apiv1.POST("/user/edituser", api.EditUserInfo)

		//获取userid
		apiv1.GET("/user/getid", api.GetUserID)

		//获取can_comment
		apiv1.GET("/user/getcomment", api.GetUserComment)

		//获取自己评论列表
		apiv1.GET("/user/getusercomment", api.UserGetComment)

		//获取消息通知
		apiv1.GET("/user/getusernoti", api.GetUserNoti)

		//删除消息通知
		apiv1.GET("/user/delusernoti", api.DeleteGetUserNoti)

		//升级会员
		apiv1.POST("/user/bevip", api.BecomeVip)

		//添加到收藏夹
		apiv1.GET("/favorite/addfavorite", v1.AddFavorite)

		//删除收藏夹指定视频
		apiv1.GET("/favorite/deletefavorite", v1.DeleteFavoriteByID)

		//获取所有标签
		apiv1.GET("/tags/getalltags", v1.GetTags)

		//获取收藏夹视频
		apiv1.GET("/favorite/getfavorite", v1.GetUserFavorite)

		//添加视频评论
		apiv1.POST("/video/addcontent", v1.AddContent)

		//获取热门视频
		apiv1.GET("/video/hotvideo", v1.GetHotVideo)

		//获取最新视频
		apiv1.GET("/video/newvideo", v1.GetNewVideo)

		//删除评论
		apiv1.GET("/video/deletecontent", v1.DeleteContent)

		//点赞评论
		apiv1.GET("/video/starcontent", v1.StarContent)

		//播放指定视频
		apiv1.GET("/video/play", v1.StreamHandler)

		//搜索视频
		apiv1.GET("/video/search", v1.SearchVideo)

		//获取所有免费视频
		apiv1.GET("/video/getallfree", v1.GetAllFreePriview)

		//获取所有VIP视频
		apiv1.GET("/video/getallvip", v1.GetAllVipPriview)

		//获取单个视频详情
		apiv1.GET("/video/getvideo", v1.GetVideoByID)

		//用户获取自己的信息
		apiv1.GET("/user/getuserinfo", api.GetUserInfo)

		//根据分类检索视频
		apiv1.GET("video/getvideobytag", v1.GetVideosByTag)
	}
	apiv2 := r.Group("api/admin")

	apiv2.Use(jwt.JWT(), verification.Verification())
	{
		//管理员删除视频
		apiv2.DELETE("/videos/deletevideo", api.AdminDeleteVideo)

		//管理员获取视频详情
		apiv2.GET("/videos/getvideobyid", v1.AdminGetVideo)

		//新建标签
		apiv2.POST("/tags/addtag", v1.AddTag)

		//管理员删除用户评论
		apiv2.GET("/comment/deletecomment", api.AdminDeleteContent)

		////删除指定标签
		apiv2.DELETE("/tags/deletetag", v1.DeleteTag)

		//新建视频
		apiv2.POST("/videos/addvideo", v1.AddVideo)

		//获取所有用户信息
		apiv2.GET("/getusers", api.AdminGetAllUser)

		//封禁用户评论功能
		apiv2.POST("/user/banusercomment", api.DisableUserComments)

		//解禁用户评论功能
		apiv2.POST("/user/restorecomment", api.AbleUserComments)

		//管理员获取所有视频信息
		apiv2.GET("/video/getallvideo", api.AdminGetAllVideoInfo)

		//管理员获取所有标签
		apiv2.GET("/tag/getalltag", api.AdminGetAllTag)

		//删除视频
		apiv2.DELETE("/video/deletevideo", v1.AdminDeleteVideo)

		//Vip视频变免费
		apiv2.GET("/video/vipvideobefree", v1.VipVideoBeFree)

		//免费视频变VIP
		apiv2.GET("/video/freevideobevip", v1.FreeVideoBeVip)
	}
	return r
}
