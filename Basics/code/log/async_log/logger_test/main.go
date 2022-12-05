package main

import (
	"time"

	"../logger"
)

var log logger.Logger

func main() {
	// log = logger.NewConsoleLogger("info")                               // 终端日志实例
	log = logger.NewFileLogger("info", "./", "testlog.log", "size", 10) // 文件日志实例，且按照大小切割
	// log = logger.NewFileLogger("info", "./", "logger_test.log", "time") // 文件日志实例，按照时间切割
	for {
		log.Trace("这是一条trace日志")
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		name := "小白"
		age := 18
		log.Error("这是一条error日志,name:%s,age:%d", name, age)
		log.Fatal("这是一条fatal日志")
		time.Sleep(time.Second * 2)
	}
}
