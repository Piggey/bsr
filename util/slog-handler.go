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

type SlogHandler struct {
	slog.Handler
	l        *log.Logger
	name     string
	addrport string
}

func NewSlogHandler(name string, addrport string, out io.Writer, slogOpts *slog.HandlerOptions) *SlogHandler {
	h := &SlogHandler{
		Handler:  slog.NewTextHandler(out, slogOpts),
		l:        log.New(out, "", 0),
		name:     name,
		addrport: addrport,
	}

	return h
}

func (sh *SlogHandler) Handle(ctx context.Context, r slog.Record) error {
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

	sh.l.Println(timeStr, sh.name, sh.addrport, level, msg, color.WhiteString(fields))
	return nil
}
