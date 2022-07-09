package mylogger

import (
	"fmt"
	"time"
)

// 终端写日志相关内容

// Logger 日志结构体
type Logger struct {
	level LogLevel
}

// NewLog 构造一个Logger
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{level: level}
}

func (l Logger) log(msg string, lv LogLevel) {
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
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, levelStr, fileName, funcName, fileNo, msg)
	}
}

func (l Logger) Trace(msg string) {
	l.log(msg, TRACE)
}

func (l Logger) Debug(msg string) {
	l.log(msg, DEBUG)
}

func (l Logger) Info(msg string) {
	l.log(msg, INFO)
}

func (l Logger) Warning(msg string) {
	l.log(msg, WARNING)
}

func (l Logger) Error(msg string) {
	l.log(msg, ERROR)
}

func (l Logger) Fatal(msg string) {
	l.log(msg, FATAL)
}
