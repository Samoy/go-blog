package service

import (
	"encoding/json"

	"github.com/samoy/go-blog/models"
	"github.com/samoy/go-blog/pkg/gredis"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/service/cacheservice"
)

// Article 文章结构体
type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageURL string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

// Add 添加文章
func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageURL,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

// Edit 编辑文章
func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageURL,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

// Get 获取文章
func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cacheservice.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}

// GetAll 获取全部文章
func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cacheservice.Article{
		TagID: a.TagID,
		State: a.State,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 3600)
	return articles, nil
}

// Delete 删除文章
func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

// ExistByID 通过ID判断文章是否存在
func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

// Count 文章计数
func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

// getMaps 获取文章映射
func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
