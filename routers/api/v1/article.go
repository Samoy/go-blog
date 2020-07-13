package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/app"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/service"
	"github.com/samoy/go-blog/pkg/setting"
	"github.com/samoy/go-blog/pkg/util"
	"github.com/unknwon/com"
)

// GetArticle 获取单篇文章
func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
	}
	articleService := service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorCheckExistArticleFail, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ErrorNotExistArticle, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorGetArticleFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, article)
}

// GetArticles 获取多篇文章
func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagID := -1
	if arg := c.PostForm("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		valid.Min(tagID, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	articleService := service.Article{
		TagID:    tagID,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorCountActicleFail, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorGetArticlesFail, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.Success, data)
}

// AddArticleForm 添加文章表单
type AddArticleForm struct {
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy     string `form:"created_by" valid:"Required;MaxSize(100)"`
	CoverImageURL string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddArticleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.Success {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := service.Tag{ID: form.TagID}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorExistTagFail, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ErrorNotExistTag, nil)
		return
	}

	articleService := service.Article{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageURL: form.CoverImageURL,
		State:         form.State,
		CreatedBy:     form.CreatedBy,
	}
	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAddArticleFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
}

// EditArticleForm 编辑文章表单
type EditArticleForm struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy    string `form:"modified_by" valid:"Required;MaxSize(100)"`
	CoverImageURL string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.Success {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := service.Article{
		ID:            form.ID,
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageURL: form.CoverImageURL,
		ModifiedBy:    form.ModifiedBy,
		State:         form.State,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorCheckExistArticleFail, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ErrorNotExistArticle, nil)
		return
	}

	tagService := service.Tag{ID: form.TagID}
	exists, err = tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorExistTagFail, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ErrorNotExistTag, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorEditArticleFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.InvalidParams, nil)
		return
	}

	articleService := service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorCheckExistArticleFail, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ErrorNotExistArticle, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorDeleteArticleFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
}
