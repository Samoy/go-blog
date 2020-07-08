package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//MySQL连接器
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/setting"
)

var db *gorm.DB

// Model 基本Model
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                             error
		dbType, dbName, user, password, host, tabPrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Fatal("Failed to get section 'database':%v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tabPrefix = sec.Key("TABLE_PREFIX").String()
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		logging.Fatal("Failed to connect database:%v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTabName string) string {
		return tabPrefix + defaultTabName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB 关闭数据库
func CloseDB() {
	defer db.Close()
}
