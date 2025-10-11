package main

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type Level = slog.Level

const (
	LevelTrace = slog.Level(-2)
)

type Logger struct {
	l   *slog.Logger
	lvl *slog.LevelVar
}

func New(level slog.Level) *Logger {
	var lvl slog.LevelVar
	lvl.Set(level)

	h := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     &lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()

				switch level {
				case LevelTrace:
					levelLabel = "trace"
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}))

	return &Logger{l: h, lvl: &lvl}
}

func (l *Logger) SetLevel(level Level) {
	l.lvl.Set(level)
}

func (l *Logger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.log(ctx, level, msg, args...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.l.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	// NOTE: 这里修改 skip 为 4，*slog.Logger.log 源码中 skip 为 3
	runtime.Callers(4, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.l.Handler().Handle(ctx, r)
}
