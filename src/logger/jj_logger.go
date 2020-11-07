/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package logger

import (
	"io"
	"log"
	"os"
)

// 日志记录器
// init()默认初始化一个文件流日志记录器
//
// 当配置为release模式时，不使用os.stdout记录日志
type Logger struct {
	// 日志的IO流
	LogWriter io.Writer
	// go标准库log
	Log       *log.Logger
	// log文件
	logFile   string
}

// 一个JJGo的日志实例
var JJGoLogger Logger

// 引入时初始化
// jjgo日志记录器，支持写文件和控制台
// 默认情况下 双写入，当文件流不可用时使用控制台输出
func InitLogger(logFile string, mode string) Logger {
	var logWriter io.Writer
	var syslog *log.Logger

	switch mode {
	case "debug":
		logFile, _ := openLogFile(logFile)
		syslog = log.New(io.MultiWriter(os.Stdout, logFile), "<JJGoLog>", log.Lshortfile|log.Ldate|log.Ltime)
		logWriter = io.MultiWriter(os.Stdout, logFile)

	case "release":
		logFile, err := openLogFile(logFile)
		if err != nil {
			syslog = log.New(os.Stdout, "<JJGoLog>", log.Lshortfile|log.Ldate|log.Ltime)
			logWriter = io.MultiWriter(os.Stdout)
		}else{
			syslog = log.New(logFile, "<JJGoLog>", log.Lshortfile|log.Ldate|log.Ltime)
			logWriter = io.MultiWriter(logFile)
		}

	default:
		syslog = log.New(os.Stdout, "<JJGoLog>", log.Lshortfile|log.Ldate|log.Ltime)
		logWriter = io.MultiWriter(os.Stdout)
	}

	JJGoLogger.LogWriter = logWriter
	JJGoLogger.Log = syslog
	JJGoLogger.logFile = logFile
	return JJGoLogger
}

// 记录错误
func (jjgolog *Logger) Error(msg string, err interface{}) {
	if err != nil {
		jjgolog.Log.Printf("Error, %s: %v\n", msg, err)
	}else {
		jjgolog.Log.Printf("Error, %s\n", msg)
	}
}

// 记录运行信息
func (jjgolog *Logger) Info(msg string) {
	jjgolog.Log.Printf("Infos: %s\n", msg)
}

// clean
func (jjgolog *Logger) Clean() {
	cleanLogFile(jjgolog.logFile)
}