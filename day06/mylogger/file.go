package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志

type FileLogger struct {
	level       LogLevel
	filePath    string   // 日志文件保存的路径
	fileName    string   // 日志文件保存的文件名
	fileObj     *os.File // 日志文件
	errFileObj  *os.File // 错误日志文件
	maxFileSize int64    // 最大日志文件大小
}

// 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) (*FileLogger, error) {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fileObj, err := os.OpenFile(path.Join(fp, fn), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open log file failed:%v", err)
		return nil, err
	}
	errFileObj, err := os.OpenFile(path.Join(fp, fn)+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error log file faied:%v", err)
		return nil, err
	}
	return &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		fileObj:     fileObj,
		errFileObj:  errFileObj,
		maxFileSize: maxSize,
	}, nil
}

func (l *FileLogger) log(msg string, lv LogLevel) {
	if lv >= l.level {
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, fileNo := getInfo(3)
		var levelStr string
		switch lv {
		case TRACE:
			levelStr = "TRACE"
		case DEBUG:
			levelStr = "DEBUG"
		case INFO:
			levelStr = "INFO"
		case WARNING:
			levelStr = "WARNING"
		case ERROR:
			levelStr = "ERROR"
		case FATAL:
			levelStr = "FATAL"
		}
		fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, levelStr, fileName, funcName, fileNo, msg)
		if lv >= ERROR {
			fmt.Fprintf(l.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, levelStr, fileName, funcName, fileNo, msg)
		}
	}
}

func (l *FileLogger) Trace(msg string) {
	l.log(msg, TRACE)
}

func (l *FileLogger) Debug(msg string) {
	l.log(msg, DEBUG)
}

func (l *FileLogger) Info(msg string) {
	l.log(msg, INFO)
}

func (l *FileLogger) Warning(msg string) {
	l.log(msg, WARNING)
}

func (l *FileLogger) Error(msg string) {
	l.log(msg, ERROR)
}

func (l *FileLogger) Fatal(msg string) {
	l.log(msg, FATAL)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
