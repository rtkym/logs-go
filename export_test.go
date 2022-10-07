package logs

func ExpSetLogLevel(s string) func() {
	tmp := envLogLevel
	envLogLevel = s

	return func() { envLogLevel = tmp }
}

func ExpSetLogFormat(s string) func() {
	tmp := envLogFormat
	envLogFormat = s

	return func() { envLogFormat = tmp }
}
