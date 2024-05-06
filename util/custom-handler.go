package util

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"time"

	"github.com/fatih/color"
)

type CustomHandler struct {
	slog.Handler
	l        *log.Logger
	name     string
	addrport string
}

func NewCustomHandler(name string, addrport string, out io.Writer, slogOpts *slog.HandlerOptions) *CustomHandler {
	h := &CustomHandler{
		Handler:  slog.NewTextHandler(out, slogOpts),
		l:        log.New(out, "", 0),
		name:     name,
		addrport: addrport,
	}

	return h
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := ""
	r.Attrs(func(a slog.Attr) bool {
		fields += fmt.Sprintf("%s=%s, ", a.Key, a.Value.String())
		return true
	})

	timeStr := r.Time.Format("[" + time.RFC3339 + "]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, h.name, h.addrport, level, msg, color.WhiteString(fields))
	return nil
}
