package service

import "github.com/samoy/go-blog/models"

// Auth 认证结构体
type Auth struct {
	Username string
	Password string
}

// Check 认证鉴权
func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
