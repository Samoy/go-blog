package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/app"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/service"
	"github.com/samoy/go-blog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth 获取认证
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	authService := service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthCheckTokenFail, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ErrorAuth, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"token": token,
	})
}
