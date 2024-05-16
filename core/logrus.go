package core

import (
	"os"

	"gvb/config"
	"gvb/global"

	"github.com/sirupsen/logrus"
)

// 初始化日志
func InitLogger(config config.Logger) *logrus.Logger {
	mLog := logrus.New()
	//日志输出位置
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(global.Conf.Logger.ShowLine) // 打印调用信息
	mLog.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	level, err := logrus.ParseLevel(global.Conf.Logger.LogLevel)
	if err != nil {
		panic(err)
	}
	mLog.SetLevel(level)
	return mLog
}
