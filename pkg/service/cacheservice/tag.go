package cacheservice

import (
	"strconv"
	"strings"

	"github.com/samoy/go-blog/pkg/e"
)

// Tag 标签缓存Model
type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

// GetTagsKey 获取多个标签缓存键
func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CacheTag,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
