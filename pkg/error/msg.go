package error

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "token超时",
	ERROR_AUTH_TOKEN:                "token生成失败",
	ERROR_AUTH:                      "token错误",
	ERROR_ADMIN_NOT_EXIST:           "该用户不存在",
	ERROR_NOT_SAME_ADMIN:            "用户名重复",
	ERROR_NOT_SEARCH:                "没有找到相关视频",
	ERROR_USER_EXIST:                "该用户名已存在，请修改",
	ERROT_EXIST_USER_DISPLAY:        "该昵称已存在，请修改后提交",
	ERROR_NOT_PHONE:                 "输入的手机号格式错误",
	ERROR_NOT_EMAIL:                 "输入的邮箱格式错误",
	ERROR_ADD_CONTENT_SUCCESS:       "添加评论成功",
	ERROR_ADD_CONTENT_ERROR:         "添加评论失败",
	ERROR_DELETE_CONTENT_NOT_MYSELF: "删除失败，不是本人的评论",
	ERROR_USER_NOT_EXIST:            "该用户不存在",
	ERROR_USER_NOT_ADMIN:            "您不是管理员，无权进行此操作",
	ERROR_FAVORITE_EXIST:            "该视频已存在您的收藏中",
	ERROR_USER_LOGIN:                "用户名或密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
