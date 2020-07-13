package export

import "github.com/samoy/go-blog/pkg/setting"

// GetExcelFullURL 获取excel url
func GetExcelFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetExcelPath() + name
}

// GetExcelPath 获取excel路径
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath 获取Excel全路径
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
