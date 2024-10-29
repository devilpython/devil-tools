package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/devilpython/devil-tools/config"
	"github.com/devilpython/devil-tools/goroutine_local"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var once sync.Once

// 获得日志实例
func GetLoggerInstance() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		//禁止logrus的输出
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("......log error:", err)
		}
		logger.Out = src
		logger.SetLevel(logrus.InfoLevel)
		conf, ok := config.GetConfigInstance()
		if ok {
			logPath := conf.LogPath
			logWriter, _ := rotatelogs.New(
				logPath+".%Y-%m-%d-%H-%M.log",
				rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
				rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
				rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
			)
			writeMap := lfshook.WriterMap{
				logrus.InfoLevel:  logWriter,
				logrus.FatalLevel: logWriter,
			}
			lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
			logger.AddHook(lfHook)
		}
	})
	return logger
}

// 记录日志信息
func Info(msg string) {
	logger.Infof(" GOROUTINE_ID[%d] | %s ", goroutine_local.GetGoroutineID(), msg)
}
