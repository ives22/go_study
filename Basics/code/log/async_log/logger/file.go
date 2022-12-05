package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

// FileLogger 往文件写日志结构体
type FileLogger struct {
	Level            LogLevel     // 日志级别
	filePath         string       // 日志文件保存的路径
	fileName         string       // 日志文件保存的文件名
	fileObj          *os.File     // info日志打开文件对象
	errFileObj       *os.File     // err以下级别打开文件对象
	splitType        string       // 切割类型，size 或者 time
	maxFileSize      int64        // 文件切割大小
	splitFileHour    int          // 记录上一次切割的小时数  针对info日志（和下面splitErrfileHour必须分开写，不然会导致只有一个日志文件被切割）
	splitErrfileHour int          // 记录上一次切割的小时数 针对err以下级别日志
	logChan          chan *logMsg // 定义一个通道
}

// logMsg 通道结构体
type logMsg struct {
	level     LogLevel
	msg       string
	fileName  string
	funcName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数：
// levelStr: 日志记录级别;
// filePath: 日志存放路径;
// fileName: 日志文件名;
// splitType: 切割类型, 如果为size表示按照大小切割(并且maxSize必须传值)，如果传入time表示按照时间切割，默认1小时切割一次;
// maxSize: 日志切割大小单位(MB)，只能传一个值。
func NewFileLogger(levelStr, filePath, fileName, splitType string, maxSize ...int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	var f1 *FileLogger
	// 如果切割类型为size，并且maxSize参数只有一个值
	if splitType == "size" && len(maxSize) == 1 {
		f1 = &FileLogger{
			Level:       logLevel,
			filePath:    filePath,
			fileName:    fileName,
			splitType:   splitType,
			maxFileSize: maxSize[0] * 1024 * 1024, // 转换为MB
			logChan:     make(chan *logMsg, 50000),
		}
	} else if splitType == "time" {
		f1 = &FileLogger{
			Level:     logLevel,
			filePath:  filePath,
			fileName:  fileName,
			splitType: splitType,
			logChan:   make(chan *logMsg, 50000),
		}
	} else {
		err := errors.New("参数错误")
		panic(err)
	}
	err = f1.initFile() // 调用initFile()方法打开文件句柄
	if err != nil {
		panic(err)
	}
	return f1
}

// initFile 打开文件对象
func (f *FileLogger) initFile() error {
	// 获取当前文件打开文件的绝对路径
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.fileObj = fileObj
	// 打开错误日志文件，这里在错误日志前面加上error，如error-xxx.log
	errFileObj, err := os.OpenFile("ERROR_"+fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.errFileObj = errFileObj
	// 第一次打开文件的时候，并将当前的小时数复制给splitHour，用于后面小时切割
	hour := time.Now().Hour()
	f.splitFileHour = hour
	f.splitErrfileHour = hour

	// 开启一个后台的goroutine去写日志
	go f.backgroundWriteLog()
	return nil
}

// enable 判断是否需要记录该日志，开关比较函数
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

// splitByFileSize 按照制定的日志大小，判断是否需要进行切割，如果大小超过设置的大小，则进行切割，如果没有超过，则不进行切割
func (f *FileLogger) splitByFileSize(file *os.File) (*os.File, error) {
	// 首先判断文件的大小，以及名字, 通过 file.Stat() 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err: %v\n", err)
		return nil, err
	}
	if fileInfo.Size() >= f.maxFileSize {
		nowStr := time.Now().Format("20060102150405")       // 获取当前时间：如20221108122342
		logName := path.Join(f.filePath, fileInfo.Name())   // 获取当前的日志文件完整路径
		newLogName := fmt.Sprintf("%s.%s", logName, nowStr) // 拼接一个日志文件备份的名字
		// 1.关闭文件
		file.Close()
		// 2.备份一下 rename   xx.log  xx.log.202211081226
		os.Rename(logName, newLogName)
		// 3.打开一个新的日志文件
		fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open new log file failed, err:%v\n", err)
			return nil, err
		}
		// 4. 将新打开的文件句柄返回
		return fileObj, nil
	}
	// 如果没有达到指定的大小，不做操作，直接返回原来文件的句柄
	return file, nil
}

// spliyByFileHour 按照小时进行切割，判断是否需要切割，如果当前时间点小时数与上一次切割时间点小时数相同则不切割，如果不同则切割
func (f *FileLogger) spliyByFileHour(file *os.File, splitHour int) (*os.File, int, error) {
	// 获取当前时间的小时数
	hour := time.Now().Hour()
	// 如果当前时间的小时数，和上一次的切割的时间小时数不相等，则进行切割
	if hour != splitHour {
		// 首先判断文件的大小，以及名字, 通过 file.Stat() 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("get file info failed, err: %v\n", err)
			return nil, 0, err
		}
		h, _ := time.ParseDuration("-1h")
		nowStr := time.Now().Add(h).Format("2006010215")    // 获取一小时前的时间：如2022110911
		logName := path.Join(f.filePath, fileInfo.Name())   // 获取当前的日志文件完整路径
		newLogName := fmt.Sprintf("%s.%s", logName, nowStr) // 拼接一个日志文件备份的名字
		// 1.关闭旧文件
		file.Close()
		// 2.备份一下 rename xx.log  xx.log.2022110911  xx.log.年月日时
		os.Rename(logName, newLogName)
		// 3.打开一个新的日志文件
		fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open new log file failed, err:%v\n", err)
			return nil, 0, err
		}
		// 4. 将新打开的文件句柄，以及新的时间点小时数返回
		return fileObj, hour, nil
	}
	// 如果当前时间点小时数和上一次时间点小时数相同，不做操作，直接返回原来文件的句柄，以及上一次的时间点小时数
	return file, splitHour, nil
}

// backgroundWriteLog 后台写入日志
func (f *FileLogger) backgroundWriteLog() {
	for {

		if f.splitType == "size" {
			// 调用 splitByFileSize 函数进行按文件大小切割相关操作
			newFile, err := f.splitByFileSize(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		} else {
			// 调用 spliyByFileHour 函数进行按时间切割
			newFile, hour, err := f.spliyByFileHour(f.fileObj, f.splitFileHour)
			if err != nil {
				return
			}
			f.fileObj = newFile    // 更新新打开的文件句柄
			f.splitFileHour = hour // 更新切割后的时间小时数
		}

		select {
		// 从通道 logChan 中取日志并进行写入文件
		case logTmp := <-f.logChan:
			// 把日志拼出来
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, unParseLogLevel(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			// 记录制定级别日志
			fmt.Fprintf(f.fileObj, logInfo)
			// 如果级别大于ERROR 级别，再将ERROR级别的日志单独记录到一个error中
			if logTmp.level >= ERROR {
				if f.splitType == "size" {
					// 调用 splitByFileSize 函数进行按文件大小切割相关操作
					newFile, err := f.splitByFileSize(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				} else {
					// 调用 spliyByFileHour 函数进行按时间切割
					newFile, hour, err := f.spliyByFileHour(f.errFileObj, f.splitErrfileHour)
					if err != nil {
						return
					}
					f.errFileObj = newFile    // 更新新打开的文件句柄
					f.splitErrfileHour = hour // 更新切割后的时间小时数
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 如果没有取到值，则睡50毫秒
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// log 记录日志的方法
func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3) // 获取调用该日志方法对应的文件名、函数名、行号
		// 先把日志发送到通道中
		// 1. 造一个logMsg对象
		logTmp := &logMsg{
			level:     level,
			msg:       msg,
			fileName:  fileName,
			funcName:  funcName,
			timestamp: timeStr,
			line:      lineNo,
		}
		// 2. 放到logChan通道中
		select {
		case f.logChan <- logTmp:
		default:
			// 如果往通道中添加时阻塞了，则把日志丢掉，保证不出现阻塞
		}

	}
}

// Trace 级别日志
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Debug 级别日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info 级别日志
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning 级别日志
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error 级别日志
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal 级别日志
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
