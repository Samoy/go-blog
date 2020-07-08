package models

import "github.com/jinzhu/gorm"

// Auth 认证Model
type Auth struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检测认证
func CheckAuth(username string, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}
