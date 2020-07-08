package setting

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/samoy/go-blog/pkg/logging"
)

var (
	// Cfg 配置文件
	Cfg *ini.File
	// RunMode 运行模式
	RunMode string
	// HTTPPort 运行端口
	HTTPPort int
	// ReadTimeout 读取超时时间
	ReadTimeout time.Duration
	// WriteTimeout 写入超时时间
	WriteTimeout time.Duration

	// PageSize 每页条数
	PageSize int
	// JwtSecret Jwt密钥
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logging.Fatalf("Failed to parse 'conf/app.ini':%v", err)
	}
	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatalf("Failed to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatalf("Failed to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("2020$06281408")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
