package util

import "github.com/samoy/go-blog/pkg/setting"

// Setup 初始化
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
