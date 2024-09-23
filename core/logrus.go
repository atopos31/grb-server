package core

import (
	"fmt"
	"io"

	"gvb/config"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// 初始化日志
func NewLogger(config config.Logger) *logrus.Logger {
	log := logrus.New()
	//日志输出位置
	var outs []io.Writer
	if config.LogTOConsole {
		outs = append(outs, log.Out)
	}

	if config.LogToFile {
		rotaW, err := rotatelogs.New(config.FilePath)
		if err != nil {
			panic(err)
		}
		outs = append(outs, rotaW)
	}

	if len(outs) > 0 {
		log.SetOutput(io.MultiWriter(outs...))
	} else {
		fmt.Println("日志未配置,默认输出到控制台")
	}

	log.SetReportCaller(config.ShowLine) // 打印调用信息
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)
	return log
}
