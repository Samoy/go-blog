package e

// MsgFlags 错误信息映射
var MsgFlags = map[int]string{
	Success:                     "成功",
	Error:                       "失败",
	InvalidParams:               "请求参数错误",
	ErrorExistTag:               "已存在该标签名称",
	ErrorNotExistTag:            "该标签不存在",
	ErrorNotExistArticle:        "该文章不存在",
	ErrorAuthCheckTokenFail:     "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:  "Token已超时",
	ErrorAuthToken:              "Token生成失败",
	ErrorAuth:                   "Token错误",
	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",
}

// GetMsg 根据code获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
