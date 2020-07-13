package service

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/samoy/go-blog/models"
	"github.com/samoy/go-blog/pkg/export"
	"github.com/samoy/go-blog/pkg/gredis"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/service/cacheservice"
	"github.com/tealeg/xlsx"
)

// Tag 标签结构体
type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

// ExistByName 通过名字判断标签是否存在
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

// ExistByID 通过ID判断标签是否存在
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

// Add 添加标签
func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

// Edit 编辑标签
func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

// Delete 删除标签
func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

// Count 标签计数
func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

// GetAll 获取全部标签
func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	cache := cacheservice.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, tags, 3600)
	return tags, nil
}

// getMaps 获取标签映射
func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}

// Export 给标签添加导出功能
func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()
	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}
	for _, v := range tags {
		values := []string{
			strconv.Itoa(int(v.ID)),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(int(v.CreatedAt.Unix())),
			v.ModifiedBy,
			strconv.Itoa(int(v.UpdatedAt.Unix())),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}
	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"
	fullPath := export.GetExcelFullPath() + filename
	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// Import 给标签添加导入功能
func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}
