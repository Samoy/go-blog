package cacheservice

import (
	"strconv"
	"strings"

	"github.com/samoy/go-blog/pkg/e"
)

// Article 文章缓存结构体
type Article struct {
	ID       int
	TagID    int
	State    int
	PageNum  int
	PageSize int
}

// GetArticleKey 获取文章缓存键
func (a *Article) GetArticleKey() string {
	return e.CacheArticle + "_" + strconv.Itoa(a.ID)
}

// GetArticlesKey 获取多个文章缓存建
func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CacheArticle,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
