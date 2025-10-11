package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", u.ID),
		slog.String("name", u.Name),
	)
}

func main() {
	// 默认日志级别为 Info
	slog.Debug("Debug")
	slog.Info("Info")
	slog.Warn("Warn")
	slog.Error("Error")

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("Set debug level")

	// 结构化日志
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))

	l.Debug("Debug struct", "test", "1")

	slog.SetDefault(l)
	log.Println("Test log print") // 同样被设置的默认 slog 影响

	// 使用强类型
	l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"info msg",
		slog.String("hello", "word"),
		slog.Int("status", 405),
	)

	// 属性分组
	l.Info(
		"info message",
		slog.Group("user", "name", "root", slog.Int("age", 20)),
	)

	// 附加自定义属性
	l = l.With("requestId", "123")
	l.Info("Test with")

	user := &User{ID: 123, Name: "test", Password: "Pass"}
	l.Info("User", "user-smg", user)
}
