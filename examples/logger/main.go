package main

import logs "github.com/rtkym/logs-go"

func main() {
	logger := logs.New()
	logger.Set("key", "val")
	logger.V("foo", "1").Info("hoge")
	logger.V("bar", "2").Info("fuga")
}
