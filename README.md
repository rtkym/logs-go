# logs-go
logs-go is a logging library wrapped around zerolog.

It sacrifices a bit of zerolog's advantage of low allocation and instead improves usability. In most cases, it can be used as-is, but in situations where performance must be pursued, it is also possible to call zerolog's API.

## Usage
### Global logger
Source:
```go
package main

import logs "github.com/rtkym/logs-go"

func main() {
	logs.Set("key", "val")
	logs.V("foo", "1").Info("hoge")
	logs.V("bar", "2").Info("fuga")
}
```
Output:
```
{"level":"info","key":"val","foo":"1","time":"2022-08-16T14:05:47.215002900+09:00","message":"hoge"}
{"level":"info","key":"val","bar":"2","time":"2022-08-16T14:05:47.263814600+09:00","message":"fuga"}
```
### Logger
Source:
```go
package main

import logs "github.com/rtkym/logs-go"

func main() {
	logger := logs.New()
	logger.Set("key", "val")
	logger.V("foo", "1").Info("hoge")
	logger.V("bar", "2").Info("fuga")
}
```
Output:
```
{"level":"info","key":"val","foo":"1","time":"2022-08-16T14:10:41.535728100+09:00","message":"hoge"}
{"level":"info","key":"val","bar":"2","time":"2022-08-16T14:10:41.575108000+09:00","message":"fuga"}
```
#### With context
```go
ctx := context.Background()
logCtx := &optctx.OptCtx{ /* some configuration options */}
ctx = optctx.NewContext(ctx, logCtx)
logger := optctx.NewLogger(ctx)
```

## Environments
### LOG_LEVEL
Supported values ​​for the environment variable `LOG_LEVEL` are `trace`, `debug`, `info`, `warning`, `error` and `fatal`. default value is `info`.
### LOG_FORMAT
Supported values ​​for the environment variable `LOG_FORMAT` are `json` and `console`. default value is `json`.

## Layout (writer)

### JSON
```go
logger := logs.New()
logger := logs.NewWithOption()
logger := logs.NewWithOption(logs.OptionWriter("json"))
logger := logs.NewWithOption(logs.OptionJSONWriter())
```

### Console
```go
logger := logs.NewWithOption(logs.OptionWriter("console"))
logger := logs.NewWithOption(logs.OptionConsoleWriter())
```

### Other
```go
buf := &bytes.Buffer{}
logger := logs.NewWithOption(func(opt *logs.Option) { opt.Writer = buf })
```
