package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/setting"
	v1 "github.com/samoy/go-blog/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//信件标签
		apiV1.POST("/tags", v1.AddTag)
		//修改标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

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
