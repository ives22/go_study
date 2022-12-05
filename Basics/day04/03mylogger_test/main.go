package main

import (
	"time"

	"../logger"
)

var log logger.Logger // 声明一个全局的接口变量

// 测试我们自己写的 日志库
func main() {
	// log = mylogger.NewConsoleLogger("Info")                               // 终端日志实例
	// log = mylogger.NewFileLogger("Info", "./", "rizhi.log", 10*1024*1024) // 文件日志实例
	// for {
	// 	log.Debug("这是一条Debug日志")
	// 	log.Info("这是一条Info日志")
	// 	log.Warning("这是一条Warning日志")
	// 	id := 10010
	// 	name := "李想"
	// 	log.Error("这是一条Error日志, id:%d, name:%s", id, name)
	// 	log.Fatal("这是一条Fatal日志")
	// 	// time.Sleep(time.Second * 1)
	// }
	// log := logger.NewConsoleLogger("erroR")                                   // 终端日志实例
	log = logger.NewFileLogger("info", "./", "logger_test.log", "size", 10) // 文件日志实例，且按照大小切割
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
