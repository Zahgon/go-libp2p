package libp2pwebrtc

import (
	"log/slog"

	logging "github.com/libp2p/go-libp2p/gologshim"
	pionLogging "github.com/pion/logging"
)

var log = logging.Logger("webrtc-transport")

var pionLog = logging.Logger("webrtc-transport-pion")

type pionLogger struct {
	*slog.Logger
}

var pLog = pionLogger{pionLog}

var _ pionLogging.LeveledLogger = pLog

func (l pionLogger) Debug(s string) { _ = "STUB: not implemented"; return }

func (l pionLogger) Debugf(s string, args ...any) { _ = "STUB: not implemented"; return }

func (l pionLogger) Error(s string) { _ = "STUB: not implemented"; return }

func (l pionLogger) Errorf(s string, args ...any) { _ = "STUB: not implemented"; return }

func (l pionLogger) Info(s string) { _ = "STUB: not implemented"; return }

func (l pionLogger) Infof(s string, args ...any) { _ = "STUB: not implemented"; return }

func (l pionLogger) Warn(s string) { _ = "STUB: not implemented"; return }

func (l pionLogger) Warnf(s string, args ...any) { _ = "STUB: not implemented"; return }

func (l pionLogger) Trace(s string) { _ = "STUB: not implemented"; return }

func (l pionLogger) Tracef(s string, args ...any) { _ = "STUB: not implemented"; return }

type loggerFactory struct{}

func (loggerFactory) NewLogger(_ string) pionLogging.LeveledLogger {
	_ = "STUB: not implemented"
	return *new(pionLogging.LeveledLogger)
}

var _ pionLogging.LoggerFactory = loggerFactory{}

var pionLoggerFactory = loggerFactory{}
