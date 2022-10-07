package logs_test

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	logs "github.com/rtkym/logs-go"
	"github.com/stretchr/testify/assert"
)

type Level int8

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return LevelTraceValue
	case DebugLevel:
		return LevelDebugValue
	case InfoLevel:
		return LevelInfoValue
	case WarnLevel:
		return LevelWarnValue
	case ErrorLevel:
		return LevelErrorValue
	case FatalLevel:
		return LevelFatalValue
	}

	return strconv.Itoa(int(l))
}

const (
	LevelTraceValue = "trace"
	LevelDebugValue = "debug"
	LevelInfoValue  = "info"
	LevelWarnValue  = "warn"
	LevelErrorValue = "error"
	LevelFatalValue = "fatal"
)

func testExec(t *testing.T, testee func(msg string), level Level, threshold Level, buf *bytes.Buffer) {
	t.Helper()

	msg := level.String() + ": " + uuid.New().String()

	buf.Reset()
	testee(msg)

	if threshold <= level {
		assert.Contains(t, buf.String(), `"`+msg+`"`)
	} else {
		assert.NotContains(t, buf.String(), `"`+msg+`"`)
	}
}

func TestGlobalLogger(t *testing.T) {
	t.Run("trace", func(t *testing.T) {
		level := TraceLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		testExec(t, logs.Trace, TraceLevel, level, buf)
		testExec(t, logs.Debug, DebugLevel, level, buf)
		testExec(t, logs.Info, InfoLevel, level, buf)
		testExec(t, logs.Warn, WarnLevel, level, buf)
		testExec(t, logs.Error, ErrorLevel, level, buf)
	})

	t.Run("debug", func(t *testing.T) {
		level := DebugLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		testExec(t, logs.Trace, TraceLevel, level, buf)
		testExec(t, logs.Debug, DebugLevel, level, buf)
		testExec(t, logs.Info, InfoLevel, level, buf)
		testExec(t, logs.Warn, WarnLevel, level, buf)
		testExec(t, logs.Error, ErrorLevel, level, buf)
	})

	t.Run("info", func(t *testing.T) {
		level := InfoLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		testExec(t, logs.Trace, TraceLevel, level, buf)
		testExec(t, logs.Debug, DebugLevel, level, buf)
		testExec(t, logs.Info, InfoLevel, level, buf)
		testExec(t, logs.Warn, WarnLevel, level, buf)
		testExec(t, logs.Error, ErrorLevel, level, buf)
	})

	t.Run("warn", func(t *testing.T) {
		level := WarnLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		testExec(t, logs.Trace, TraceLevel, level, buf)
		testExec(t, logs.Debug, DebugLevel, level, buf)
		testExec(t, logs.Info, InfoLevel, level, buf)
		testExec(t, logs.Warn, WarnLevel, level, buf)
		testExec(t, logs.Error, ErrorLevel, level, buf)
	})

	t.Run("error", func(t *testing.T) {
		level := ErrorLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		testExec(t, logs.Trace, TraceLevel, level, buf)
		testExec(t, logs.Debug, DebugLevel, level, buf)
		testExec(t, logs.Info, InfoLevel, level, buf)
		testExec(t, logs.Warn, WarnLevel, level, buf)
		testExec(t, logs.Error, ErrorLevel, level, buf)
	})

	t.Run("set,with", func(t *testing.T) {
		level := InfoLevel

		buf := &bytes.Buffer{}
		logs.GlobalLoggerOptions = []logs.OptionFunc{logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf }}
		logs.InitGlobalLogger()

		logs.Set("set1", "a")
		logs.Set("set2", "b")

		buf.Reset()
		logs.Entry().V("with1", "1").V("with2", "2").Info("test msg1")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"with1":"1"`)
		assert.Contains(t, buf.String(), `"with2":"2"`)
		assert.Contains(t, buf.String(), `"test msg1"`)

		buf.Reset()
		logs.V("with1", "X").V("with2", "Y").Info("test msg2")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"with1":"X"`)
		assert.Contains(t, buf.String(), `"with2":"Y"`)
		assert.Contains(t, buf.String(), `"test msg2"`)

		buf.Reset()
		logs.E(errors.New("test error")).Info("test msg3")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"error":{"message":"test error"}`)
		assert.Contains(t, buf.String(), `"test msg3"`)
	})
}

func TestLogger(t *testing.T) {
	t.Run("trace", func(t *testing.T) {
		level := TraceLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		testExec(t, logger.Trace, TraceLevel, level, buf)
		testExec(t, logger.Debug, DebugLevel, level, buf)
		testExec(t, logger.Info, InfoLevel, level, buf)
		testExec(t, logger.Warn, WarnLevel, level, buf)
		testExec(t, logger.Error, ErrorLevel, level, buf)
	})

	t.Run("debug", func(t *testing.T) {
		level := DebugLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		testExec(t, logger.Trace, TraceLevel, level, buf)
		testExec(t, logger.Debug, DebugLevel, level, buf)
		testExec(t, logger.Info, InfoLevel, level, buf)
		testExec(t, logger.Warn, WarnLevel, level, buf)
		testExec(t, logger.Error, ErrorLevel, level, buf)
	})

	t.Run("info", func(t *testing.T) {
		level := InfoLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		testExec(t, logger.Trace, TraceLevel, level, buf)
		testExec(t, logger.Debug, DebugLevel, level, buf)
		testExec(t, logger.Info, InfoLevel, level, buf)
		testExec(t, logger.Warn, WarnLevel, level, buf)
		testExec(t, logger.Error, ErrorLevel, level, buf)
	})

	t.Run("warn", func(t *testing.T) {
		level := WarnLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		testExec(t, logger.Trace, TraceLevel, level, buf)
		testExec(t, logger.Debug, DebugLevel, level, buf)
		testExec(t, logger.Info, InfoLevel, level, buf)
		testExec(t, logger.Warn, WarnLevel, level, buf)
		testExec(t, logger.Error, ErrorLevel, level, buf)
	})

	t.Run("error", func(t *testing.T) {
		level := ErrorLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		testExec(t, logger.Trace, TraceLevel, level, buf)
		testExec(t, logger.Debug, DebugLevel, level, buf)
		testExec(t, logger.Info, InfoLevel, level, buf)
		testExec(t, logger.Warn, WarnLevel, level, buf)
		testExec(t, logger.Error, ErrorLevel, level, buf)
	})

	t.Run("set,with", func(t *testing.T) {
		level := InfoLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		logger.Set("set1", "a")
		logger.Set("set2", "b")

		buf.Reset()
		logger.Entry().V("with1", "1").V("with2", "2").Info("test msg1")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"with1":"1"`)
		assert.Contains(t, buf.String(), `"with2":"2"`)
		assert.Contains(t, buf.String(), `"test msg1"`)

		buf.Reset()
		logger.V("with1", "X").V("with2", "Y").Info("test msg2")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"with1":"X"`)
		assert.Contains(t, buf.String(), `"with2":"Y"`)
		assert.Contains(t, buf.String(), `"test msg2"`)

		buf.Reset()
		logger.E(errors.New("test error")).Info("test msg3")

		assert.Contains(t, buf.String(), `"set1":"a"`)
		assert.Contains(t, buf.String(), `"set2":"b"`)
		assert.Contains(t, buf.String(), `"error":{"message":"test error"}`)
		assert.Contains(t, buf.String(), `"test msg3"`)
	})

	t.Run("WithError", func(t *testing.T) {
		level := InfoLevel

		buf := &bytes.Buffer{}
		logger := logs.NewWithOption(logs.OptionLevel(level.String()), func(opt *logs.Option) { opt.Writer = buf })

		buf.Reset()
		logger.E(&MarshalableError{}).Info("test msg1")

		assert.Contains(t, buf.String(), `"error":{"key":"value"}`)
		assert.Contains(t, buf.String(), `"test msg1"`)
	})
}

type MarshalableError struct{}

func (err *MarshalableError) Error() string {
	return "MarshalableError"
}

func (err *MarshalableError) MarshalJSON() ([]byte, error) {
	return []byte(`{"key":"value"}`), nil
}

type ValuesError1 struct{}

func (err *ValuesError1) Values() map[string]interface{} {
	return map[string]interface{}{"vKey": "vValue"}
}

type ValuesError2 struct{}

func (err *ValuesError2) Values() []interface{} {
	return []interface{}{"hoge"}
}

type ValuesError3 struct{}

func (err *ValuesError3) Values() []string {
	return []string{"fuga"}
}

func TestConsoleLogger(t *testing.T) {
	var cWriter *zerolog.ConsoleWriter

	logger := logs.NewWithOption(func(opt *logs.Option) {
		logs.OptionConsoleWriter()(opt)

		if writer, ok := opt.Writer.(*zerolog.ConsoleWriter); ok {
			cWriter = writer
		} else {
			panic("illegal writer instance: " + fmt.Sprintf("%T", opt.Writer))
		}
	})

	buf := &bytes.Buffer{}
	cWriter.Out = buf

	logger.Set("set1", "a")
	logger.Set("set2", "b")

	buf.Reset()
	logger.Info("test masg")

	assert.Contains(t, buf.String(), `info  test masg`)
	assert.Contains(t, buf.String(), `{set1:a}`)
	assert.Contains(t, buf.String(), `{set2:b}`)

	buf.Reset()
	logger.Error("test masg")

	assert.Contains(t, buf.String(), `error test masg`)
	assert.Contains(t, buf.String(), `{set1:a}`)
	assert.Contains(t, buf.String(), `{set2:b}`)
}

func TestNewConsole(t *testing.T) {
	defer logs.ExpSetLogFormat("console")()

	logger := logs.New()

	assert.NotNil(t, logger)
}

func TestNewJSON(t *testing.T) {
	defer logs.ExpSetLogFormat("json")()

	logger := logs.New()

	assert.NotNil(t, logger)
}
