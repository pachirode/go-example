package main

import (
	"context"
	"io"
	"log/slog"
)

type Handler struct {
	slog.Handler
}

func NewHandler(w io.Writer, opts *slog.HandlerOptions) *Handler {
	return &Handler{
		Handler: slog.NewJSONHandler(w, opts),
	}
}

// Enable 当前日志级别是否开启
func (h *Handler) Enable(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

// Handle 处理日志记录
func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	record.Add("custom", "handler")
	return h.Handler.Handle(ctx, record)
}
