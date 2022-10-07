package logs

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	envLogLevel  = os.Getenv("LOG_LEVEL")  // nolint:gochecknoglobals
	envLogFormat = os.Getenv("LOG_FORMAT") // nolint:gochecknoglobals
)

// New returns a new Logger.
func New() *Logger {
	return NewWithOption(
		OptionLevel(envLogLevel),
		OptionWriter(envLogFormat),
	)
}

// NewWithOption returns a new Logger with options.
func NewWithOption(opts ...OptionFunc) *Logger {
	opt := &Option{
		Level:  zerolog.InfoLevel,
		Writer: os.Stdout,
	}

	for _, fn := range opts {
		fn(opt)
	}

	logger := zerolog.New(opt.Writer).Level(opt.Level).With().Timestamp().Logger()

	return &Logger{
		zeroLogger: logger,
	}
}

type Option struct {
	Level  zerolog.Level
	Writer io.Writer
}

type OptionFunc func(opt *Option)
