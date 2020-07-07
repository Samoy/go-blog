package api

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/models"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth 获取认证
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.InvalidParams
	var msg string
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token
				code = e.Success
			}
		} else {
			code = e.ErrorAuth
		}
		msg = e.GetMsg(code)
	} else {
		for _, err := range valid.Errors {
			msg += err.Message + ";"
		}
		msg = strings.TrimRight(msg, ";")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
