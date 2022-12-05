package logger

import (
	"fmt"
	"time"
)

// 往终端写日志

// ConsoleLogger 往终端输出日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: logLevel,
	}
}

// 开关比较函数
func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if c.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", timeStr, unParseLogLevel(level), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
