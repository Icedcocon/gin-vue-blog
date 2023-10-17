package utils

import (
	"backend/pkg/config"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// 全局日志指针
var GLogger *logrus.Logger

// 自定义日志格式
type MyFormatter struct{}

// 实现Formatter接口
func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取日志缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 设置日志格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	// 判断是否有函数调用信息
	if entry.HasCaller() {
		// 获取文件名
		fName := filepath.Base(entry.Caller.File)
		// 格式化日志
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level.String(), fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		// 格式化日志
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level.String(), entry.Message)
	}

	// 写入缓冲区
	b.WriteString(newLog)
	return b.Bytes(), nil
}

// 初始化日志
func InitGLogger() {
	// GLogger = logrus.New()
	GLogger = logrus.StandardLogger()
	level, err := logrus.ParseLevel(config.GlobalConfig.LOG.Level)
	if err != nil {
		level = logrus.TraceLevel
	}
	// 设置日志级别
	GLogger.SetLevel(level)
	// 返回函数调用的文件名和行号
	GLogger.SetReportCaller(true)
	// 设置日志输出
	f, _ := os.OpenFile(config.GlobalConfig.LOG.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	GLogger.SetOutput(io.MultiWriter(os.Stdout, f))
	// 设置日志格式
	// GLogger.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors:     true,
	// 	FullTimestamp:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	GLogger.SetFormatter(&MyFormatter{})
	// 设置日志不加锁
	GLogger.SetNoLock()
}

// 自定义日志输出时间格式
func SetGLoggerFormatter(formatter logrus.Formatter) {
	GLogger.SetFormatter(formatter)
}
