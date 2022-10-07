package logs

import (
	"github.com/rs/zerolog"
)

// LogEntry is one record of logging.
type LogEntry struct {
	logger *Logger
	values map[string]interface{}
}

func (x *LogEntry) bind(ev *zerolog.Event) {
	for k, v := range x.values {
		ev.Interface(k, v)
	}
}

// Trace outputs messages at trace level.
func (x *LogEntry) Trace(msg string) {
	ev := x.logger.zeroLogger.Trace()
	x.bind(ev)
	ev.Msg(msg)
}

// Debug outputs messages at debug level.
func (x *LogEntry) Debug(msg string) {
	ev := x.logger.zeroLogger.Debug()
	x.bind(ev)
	ev.Msg(msg)
}

// Info outputs messages at info level.
func (x *LogEntry) Info(msg string) {
	ev := x.logger.zeroLogger.Info()
	x.bind(ev)
	ev.Msg(msg)
}

// Warn outputs messages at warn level.
func (x *LogEntry) Warn(msg string) {
	ev := x.logger.zeroLogger.Warn()
	x.bind(ev)
	ev.Msg(msg)
}

// Error outputs messages at error level.
func (x *LogEntry) Error(msg string) {
	ev := x.logger.zeroLogger.Error()
	x.bind(ev)
	ev.Msg(msg)
}

// Fatal outputs messages at fatal level.
func (x *LogEntry) Fatal(msg string) {
	ev := x.logger.zeroLogger.Fatal()
	x.bind(ev)
	ev.Msg(msg)
}

// V adds key and value attribute to log message.
func (x *LogEntry) V(key string, value interface{}) *LogEntry {
	x.values[key] = value
	return x
}

// E adds error attribute to log message.
func (x *LogEntry) E(err error) *LogEntry {
	if _, ok := err.(interface{ MarshalJSON() ([]byte, error) }); ok { // nolint:errorlint
		return x.V("error", err)
	}

	return x.V("error", map[string]string{"message": err.Error()})
}
