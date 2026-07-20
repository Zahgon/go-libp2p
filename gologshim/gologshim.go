package gologshim

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

var lvlToLower = map[slog.Level]slog.Value{
	slog.LevelDebug: slog.StringValue("debug"),
	slog.LevelInfo:  slog.StringValue("info"),
	slog.LevelWarn:  slog.StringValue("warn"),
	slog.LevelError: slog.StringValue("error"),
}

var defaultHandler atomic.Pointer[slog.Handler]

func SetDefaultHandler(handler slog.Handler) { _ = "STUB: not implemented"; return }

type dynamicHandler struct {
	system  string
	config  *Config
	once    sync.Once
	handler slog.Handler
}

func (h *dynamicHandler) ensureHandler() slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *dynamicHandler) createFallbackHandler() slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *dynamicHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *dynamicHandler) Handle(ctx context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *dynamicHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *dynamicHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func Logger(system string) *slog.Logger { _ = "STUB: not implemented"; return nil }

type logFormat = int

const (
	logFormatText logFormat = iota
	logFormatJSON
)

type Config struct {
	fallbackLvl   slog.Level
	systemToLevel map[string]slog.Level
	format        logFormat
	addSource     bool
	labels        []slog.Attr
}

func (c *Config) LevelForSystem(system string) slog.Level {
	_ = "STUB: not implemented"
	return *new(slog.Level)
}

var ConfigFromEnv func() *Config = sync.OnceValue(func() *Config {
	fallback, systemToLevel, err := parseIPFSGoLogEnv(os.Getenv("GOLOG_LOG_LEVEL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse GOLOG_LOG_LEVEL: %v", err)
		fallback = slog.LevelInfo
	}
	c := &Config{
		fallbackLvl:   fallback,
		systemToLevel: systemToLevel,
		addSource:     true,
	}

	logFmt := os.Getenv("GOLOG_LOG_FORMAT")
	if logFmt == "" {
		logFmt = os.Getenv("GOLOG_LOG_FMT")
	}
	if logFmt == "json" {
		c.format = logFormatJSON
	}

	logFmt = os.Getenv("GOLOG_LOG_ADD_SOURCE")
	if logFmt == "0" || logFmt == "false" {
		c.addSource = false
	}

	labels := os.Getenv("GOLOG_LOG_LABELS")
	if labels != "" {
		labels := strings.Split(labels, ",")
		if len(labels) > 0 {
			for _, label := range labels {
				kv := strings.SplitN(label, "=", 2)
				if len(kv) == 2 {
					c.labels = append(c.labels, slog.String(kv[0], kv[1]))
				} else {
					fmt.Fprintf(os.Stderr, "Invalid label format: %s", label)
				}
			}
		}
	}

	return c
})

func parseIPFSGoLogEnv(loggingLevelEnvStr string) (slog.Level, map[string]slog.Level, error) {
	_ = "STUB: not implemented"
	return *new(slog.Level), nil, nil
}
