package util

import (
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/setting"
	"github.com/unknwon/com"
)

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
