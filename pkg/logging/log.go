package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level 日志级别
type Level int

var (
	// F 文件
	F *os.File
	// DefaultPrefix 默认前缀
	DefaultPrefix = ""
	// DefaultCallerDepth 默认调用深度
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	// DEBUG 调试级别，最详细的日志
	DEBUG Level = iota
	// INFO 信息级别
	INFO
	// WARNING 警告级别
	WARNING
	// ERROR 错误级别
	ERROR
	// FATAL 致命级别
	FATAL
)

// Setup 日志初始化
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug 打印调试日志
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Debugf 格式化打印调试日志
func Debugf(format string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format, v...)
}

// Info 打印信息日志
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Infof 格式化打印信息日志
func Infof(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v...)
}

// Warn 打印警告日志
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Warnf 格式化打印警告日志
func Warnf(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v...)
}

// Error 打印错误日志
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Errorf 格式化打印错误日志
func Errorf(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v...)
}

// Fatal 打印致命日志
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// Fatalf 格式化打印致命日志
func Fatalf(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(format, v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
