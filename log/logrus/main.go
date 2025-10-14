package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func base() {
	logrus.Debug("debug msg")
	logrus.Info("Info msg")
	// 输出格式 logfmt 风格
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.Warn("Warn msg")

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Trace msg")
	logrus.Debug("Debug msg")
	logrus.Info("Info msg")
	logrus.Warn("Warn msg")
	logrus.Warning("Warning msg")
}

func customLogger() {
	log := logrus.New()

	// 同时写到多个输出
	w1 := os.Stdout
	w2, _ := os.OpenFile("demo.log", os.O_WRONLY|os.O_CREATE, 0644)
	log.SetOutput(io.MultiWriter(w1, w2))

	log.SetFormatter(&logrus.JSONFormatter{})
	log.Formatter = &logrus.TextFormatter{
		// 关闭控制台颜色输出
		DisableColors: true,
		// 允许用户自定义默认字段名
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "message",
		},
	}

	logger := log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	})
	logger.Warn("Warn msg")
	logger.Error("Error msg")
}

func main() {
	//base()
	setHook()
	customLogger()
}
