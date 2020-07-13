package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/samoy/go-blog/pkg/logging"
)

// MarkErrors 给出错误信息
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
