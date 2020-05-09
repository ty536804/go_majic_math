package Mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//日志级别
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

//往终端写日志相关内容
//日志结构体
type Logger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARING":
		return WARING, nil
	case "FATAL":
		return FATAL, nil
	case "ERROR":
		return ERROR, nil
	default:
		return UNKNOWN, errors.New("无效的日志级别")
	}
}

//构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{Level: level}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now().Format("2006-01-02 15:03:04")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now().Format("2006-01-02 15:03:04")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l Logger) Waring(msg string) {
	if l.enable(WARING) {
		now := time.Now().Format("2006-01-02 15:03:04")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now().Format("2006-01-02 15:03:04")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now().Format("2006-01-02 15:03:04")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}
