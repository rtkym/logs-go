package main

import logs "github.com/rtkym/logs-go"

func main() {
	logs.Set("key", "val")
	logs.V("foo", "1").Info("hoge")
	logs.V("bar", "2").Info("fuga")
}
