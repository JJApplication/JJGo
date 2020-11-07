/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package logger

import (
	"io/ioutil"
	"os"
)

// 获取log文件路径 打开日志流
// 并没有保证流被正常关闭
func openLogFile(logFile string) (*os.File, error) {
	filePath := logFile
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	return file, err
}

// 通过overwrite重写日志达到清空的目的
func cleanLogFile(logFile string) {
	filePath := logFile
	err := ioutil.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		JJGoLogger.Error("Clean Log failed", err)
	}
}