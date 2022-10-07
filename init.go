package rtlog

import "github.com/rs/zerolog"

const DefaultDatetimeFormat = "2006-01-02T15:04:05.000000000Z07:00"

// init initializes the rtlog package
func init() { //nolint:gochecknoinits
	zerolog.TimeFieldFormat = DefaultDatetimeFormat

	InitGlobalLogger()
}
