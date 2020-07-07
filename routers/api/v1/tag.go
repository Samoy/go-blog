package v1

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/models"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/setting"
	"github.com/samoy/go-blog/pkg/util"
	"github.com/unknwon/com"
)

// GetTags 获取标签列表
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	state := 1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.Success
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag 添加标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.InvalidParams
	var msg string
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.Success
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ErrorExistTag
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
		"msg":  msg,
		"data": make(map[string]string),
	})
}

// EditTag 编辑标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.InvalidParams
	var msg string
	if !valid.HasErrors() {
		code = e.Success
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ErrorNotExistTag
		}
		msg = e.GetMsg(code)
	} else {
		for _, err := range valid.Errors {
			print(err.Message)
			msg += err.Message + ";"
		}
		msg = strings.TrimRight(msg, ";")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	var msg string
	if !valid.HasErrors() {
		code = e.Success
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ErrorNotExistTag
		}
		msg = e.GetMsg(code)
	} else {
		for _, err := range valid.Errors {
			print(err.Message)
			msg += err.Message + ";"
		}
		msg = strings.TrimRight(msg, ";")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})
}
