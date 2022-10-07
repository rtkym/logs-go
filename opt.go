package logs

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// OptionLevel returns an OptionFunc for configuring log levels.
func OptionLevel(level string) OptionFunc {
	return func(opt *Option) {
		var zerologLevel zerolog.Level

		switch strings.ToLower(level) {
		case "trace":
			zerologLevel = zerolog.TraceLevel
		case "debug":
			zerologLevel = zerolog.DebugLevel
		case "info":
			zerologLevel = zerolog.InfoLevel
		case "warn":
			zerologLevel = zerolog.WarnLevel
		case "error":
			zerologLevel = zerolog.ErrorLevel
		}

		opt.Level = zerologLevel
	}
}

// OptionWriter returns an OptionFunc for configuring log format.
func OptionWriter(format string) OptionFunc {
	return func(opt *Option) {
		switch strings.ToLower(envLogFormat) {
		case "console":
			OptionConsoleWriter()(opt)
		case "json":
			OptionJSONWriter()(opt)
		}
	}
}

// OptionJSONWriter returns an OptionFunc for configuring json format.
func OptionJSONWriter() OptionFunc {
	return func(opt *Option) {
		opt.Writer = os.Stdout
	}
}

// OptionConsoleWriter returns an OptionFunc for configuring console format.
func OptionConsoleWriter() OptionFunc {
	return func(opt *Option) {
		writer := &zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat}
		writer.FormatLevel = func(i interface{}) string {
			return fmt.Sprintf("%-5s", i)
		}
		writer.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		writer.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("{%s:", i)
		}
		writer.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%v}", i)
		}

		opt.Writer = writer
	}
}
