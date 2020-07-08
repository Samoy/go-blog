package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/util"
)

// JWT JWT中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.Success
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				logging.Error("token", err)
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
