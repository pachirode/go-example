package main

import (
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
)

type emailHook struct {
}

// Levels 所有日志级别生效
func (hook *emailHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *emailHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = "email"

	msg, _ := entry.String()
	fakeSendEmail(msg)
	return nil
}
func fakeSendEmail(msg string) {
	println("Send msg: ", msg)
}

func newRotateHook() logrus.Hook {
	hook, _ := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   "general.log",
			MaxSize:    100,
			MaxBackups: 1,
			MaxAge:     1,
			Compress:   false,
			LocalTime:  false,
		},
		logrus.InfoLevel,
		&logrus.TextFormatter{DisableColors: true},
		&lumberjackrus.LogFileOpts{ // 针对不同日志级别的配置
			logrus.TraceLevel: &lumberjackrus.LogFile{
				Filename: "trace.log",
			},
			logrus.ErrorLevel: &lumberjackrus.LogFile{
				Filename:   "error.log",
				MaxSize:    10,    // 日志文件在轮转之前的最大大小，默认 100 MB
				MaxBackups: 10,    // 保留旧日志文件的最大数量
				MaxAge:     10,    // 保留旧日志文件的最大天数
				Compress:   true,  // 是否使用 gzip 对日志文件进行压缩归档
				LocalTime:  false, // 是否使用本地时间，默认 UTC 时间
			},
		},
	)

	return hook
}

func setHook() {
	logrus.AddHook(&emailHook{})
	logrus.AddHook(newRotateHook())
}
