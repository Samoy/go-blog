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

// Setup 数据库初始化
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		logging.Fatalf("Failed to connect database:%v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTabName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTabName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&Tag{}, &Auth{}, &Article{})
}

// CloseDB 关闭数据库
func CloseDB() {
	defer db.Close()
}
