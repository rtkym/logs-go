package rtlog

import "github.com/rs/zerolog"

var (
	// gLogger is global logger instance
	gLogger *Logger // nolint:gochecknoglobals

	// GlobalLoggerOptions is gLogger configurations options.
	GlobalLoggerOptions []OptionFunc // nolint:gochecknoglobals
)

// InitGlobalLogger initialize global logger
func InitGlobalLogger() {
	if len(GlobalLoggerOptions) != 0 {
		gLogger = NewWithOption(GlobalLoggerOptions...)
	} else {
		gLogger = New()
	}
}

// Entry returns a new LogEntry
func Entry() *LogEntry { return gLogger.Entry() }

// Trace outputs messages at trace level.
func Trace(msg string) { gLogger.Entry().Trace(msg) }

// Debug outputs messages at debug level.
func Debug(msg string) { gLogger.Entry().Debug(msg) }

// Info outputs messages at info level.
func Info(msg string) { gLogger.Entry().Info(msg) }

// Warn outputs messages at warn level.
func Warn(msg string) { gLogger.Entry().Warn(msg) }

// Error outputs messages at error level.
func Error(msg string) { gLogger.Entry().Error(msg) }

// Fatal outputs messages at fatal level.
func Fatal(msg string) { gLogger.Entry().Fatal(msg) }

// V adds key and value attribute to log message.
func V(key string, value interface{}) *LogEntry {
	return gLogger.Entry().V(key, value)
}

// E adds error attribute to log message.
func E(err error) *LogEntry {
	return gLogger.Entry().E(err)
}

// Set saves key and value to logger. The key and value are output permanently
func Set(key string, value interface{}) {
	gLogger.Set(key, value)
}

// SetWithZC saves key and value to logger. The key and value are output permanently
func SetWithZC(fn func(zc ZC) ZC) {
	gLogger.SetWithZC(fn)
}

// With gets zerolog.Context
func With() zerolog.Context {
	return gLogger.With()
}

// ZC is alias for zerolog.Context.
type ZC = zerolog.Context
