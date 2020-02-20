package error

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token超时",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "token错误",
	ERROR_ADMIN_NOT_EXIST:          "该用户不存在",
	ERROR_NOT_SAME_ADMIN:           "用户名重复",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
