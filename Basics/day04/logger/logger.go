package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 基于uint16造一个Loglevel类型
type LogLevel uint16

// Logger 接口
type Logger interface {
	Trace(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// 定义常量，日志级别用于比较使用
const (
	UNKNOWN LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// parseLogLevel 将传入的级别转换为LogLevel类型
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s) // 将传入的字符串全部转为大写
	switch s {
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的级别类型")
		return UNKNOWN, err
	}
}

// unParseLogLevel 将传入的对应级别的LogLevel类型转换为字符串类型
func unParseLogLevel(level LogLevel) string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// getInfo 获取程序当前执行的信息
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return
}
