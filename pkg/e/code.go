package e

const (
	// Success 成功
	Success = 200
	// Error 失败
	Error = 500
	// InvalidParams 无效的参数
	InvalidParams = 400

	// ErrorExistTag 标签已存在
	ErrorExistTag = 10001
	// ErrorNotExistTag 标签不存在
	ErrorNotExistTag = 10002
	// ErrorNotExistArticle 文章不存在
	ErrorNotExistArticle = 10003

	// ErrorAuthCheckTokenFail Token鉴权失败
	ErrorAuthCheckTokenFail = 20001
	// ErrorAuthCheckTokenTimeout Token已超时
	ErrorAuthCheckTokenTimeout = 20002
	// ErrorAuthToken  Token生成失败
	ErrorAuthToken = 20003
	// ErrorAuth Token错误
	ErrorAuth = 20004
	// ErrorUploadSaveImageFail 保存图片失败
	ErrorUploadSaveImageFail = 30001
	// ErrorUploadCheckImageFail 检查图片失败
	ErrorUploadCheckImageFail = 30002
	// ErrorUploadCheckImageFormat 校验图片错误，图片格式或大小有问题
	ErrorUploadCheckImageFormat = 30003
)
