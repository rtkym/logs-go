package logs

import (
	"github.com/rs/zerolog"
)

// Logger provides basic logging functionality.
type Logger struct {
	zeroLogger zerolog.Logger
}

// Entry returns a new LogEntry
func (x *Logger) Entry() *LogEntry {
	return &LogEntry{
		logger: x,
		values: make(map[string]interface{}),
	}
}

// Trace outputs messages at trace level.
func (x *Logger) Trace(msg string) { x.Entry().Trace(msg) }

// Debug outputs messages at debug level.
func (x *Logger) Debug(msg string) { x.Entry().Debug(msg) }

// Info outputs messages at info level.
func (x *Logger) Info(msg string) { x.Entry().Info(msg) }

// Warn outputs messages at warn level.
func (x *Logger) Warn(msg string) { x.Entry().Warn(msg) }

// Error outputs messages at error level.
func (x *Logger) Error(msg string) { x.Entry().Error(msg) }

// Fatal outputs messages at fatal level.
func (x *Logger) Fatal(msg string) { x.Entry().Fatal(msg) }

// V adds key and value attribute to log message.
func (x *Logger) V(key string, value interface{}) *LogEntry {
	return x.Entry().V(key, value)
}

// E adds error attribute to log message.
func (x *Logger) E(err error) *LogEntry {
	return x.Entry().E(err)
}

// Set saves key and value attribute to logger. The attribute are output permanently.
func (x *Logger) Set(key string, value interface{}) {
	x.SetWithZC(func(zc ZC) ZC { return zc.Interface(key, value) })
}

// SetWithZC saves key and value to logger. The key and value are output permanently
func (x *Logger) SetWithZC(fn func(zc ZC) ZC) {
	x.zeroLogger = fn(x.With()).Logger()
}

// With gets zerolog.Context
func (x *Logger) With() zerolog.Context {
	return x.zeroLogger.With()
}
