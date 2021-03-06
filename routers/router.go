package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/middleware/jwt"
	"github.com/samoy/go-blog/pkg/export"
	"github.com/samoy/go-blog/pkg/setting"
	"github.com/samoy/go-blog/pkg/upload"
	"github.com/samoy/go-blog/routers/api"
	v1 "github.com/samoy/go-blog/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新增标签
		apiV1.POST("/tags", v1.AddTag)
		//修改标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
