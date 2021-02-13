/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package logger

import (
	"fmt"
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

// linux支持的16进制颜色
const colorString = "\x1b[0;%dm%s\x1b[0m"
const(
	RED = 31
	GREEN = 32
	YELLOW = 33
	PURPLE = 35
	CYAN = 36
	WHITE = 37
)

var colorConf = ColorConfig()

func Red(str string) string {
	// 默认的日志是不打印颜色的 为了使用颜色日志 需要开始ini的color选项
	if colorConf {
		return fmt.Sprintf(colorString, RED, str)
	}
	return str
}

func Green(str string) string {
	if colorConf {
		return fmt.Sprintf(colorString, GREEN, str)
	}
	return str
}

func Purple(str string) string {
	if colorConf {
		return fmt.Sprintf(colorString, PURPLE, str)
	}
	return str
}

func Yellow(str string) string {
	if colorConf {
		return fmt.Sprintf(colorString, YELLOW, str)
	}
	return str
}

func Cyan(str string) string {
	if colorConf {
		return fmt.Sprintf(colorString, CYAN, str)
	}
	return str
}

func White(str string) string {
	if colorConf {
		return fmt.Sprintf(colorString, WHITE, str)
	}
	return str
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
		syslog = log.New(io.MultiWriter(os.Stdout, logFile), Purple("<JJGoLog>"), log.Lshortfile|log.Ldate|log.Ltime)
		logWriter = io.MultiWriter(os.Stdout, logFile)

	case "release":
		logFile, err := openLogFile(logFile)
		if err != nil {
			syslog = log.New(os.Stdout, Purple("<JJGoLog>"), log.Lshortfile|log.Ldate|log.Ltime)
			logWriter = io.MultiWriter(os.Stdout)
		}else{
			syslog = log.New(logFile, Purple("<JJGoLog>"), log.Lshortfile|log.Ldate|log.Ltime)
			logWriter = io.MultiWriter(logFile)
		}

	default:
		syslog = log.New(os.Stdout, Purple("<JJGoLog>"), log.Lshortfile|log.Ldate|log.Ltime)
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
		jjgolog.Log.Printf(Red("[Error]") + " %s %v\n", msg, err)
	}else {
		jjgolog.Log.Printf(Red("[Error]") + " %s\n", msg)
	}
}

// 记录运行信息
func (jjgolog *Logger) Info(msg string) {
	jjgolog.Log.Printf(Cyan("[Info]") + " %s\n", msg)
}

// 记录运行警告信息
func (jjgolog *Logger) Warning(msg string, args ...interface{}) {
	if len(args) >= 0 {
		jjgolog.Log.Printf(Yellow("[Warning]") + " %s %v\n", msg, args)
	}else {
		jjgolog.Log.Printf(Yellow("[Warning]") + " %s\n", msg)
	}
}

// clean
func (jjgolog *Logger) Clean() {
	cleanLogFile(jjgolog.logFile)
}