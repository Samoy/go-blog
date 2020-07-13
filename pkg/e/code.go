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
	// ErrorCheckExistArticleFail 检测文章失败
	ErrorCheckExistArticleFail = 10004
	// ErrorGetArticleFail 获取文章失败
	ErrorGetArticleFail = 10005
	// ErrorCountActicleFail 获取文章总数失败
	ErrorCountActicleFail = 10006
	// ErrorGetArticlesFail 获取多篇文章失败
	ErrorGetArticlesFail = 10007
	// ErrorExistTagFail 检测标签失败
	ErrorExistTagFail = 10008
	// ErrorAddArticleFail 添加文章失败
	ErrorAddArticleFail = 10009
	// ErrorEditArticleFail 编辑文章失败
	ErrorEditArticleFail = 10010
	// ErrorDeleteArticleFail 删除文章失败
	ErrorDeleteArticleFail = 10011
	// ErrorGetTagsFail 获取标签失败
	ErrorGetTagsFail = 10012
	// ErrorCountTagFail 获取标签总数失败
	ErrorCountTagFail = 10013
	// ErrorAddTagFail 添加标签失败
	ErrorAddTagFail = 10014
	// ErrorEditTagFail 编辑标签失败
	ErrorEditTagFail = 10015
	// ErrorDeleteTagFail 删除标签失败
	ErrorDeleteTagFail = 10016
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
