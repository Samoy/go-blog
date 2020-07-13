package app

import (
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/e"
)

// Gin Gin结构体
type Gin struct {
	C *gin.Context
}

// Response 给出响应信息
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
