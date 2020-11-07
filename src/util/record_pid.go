/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"jjgo/src/logger"
)

// 记录运行的PID
// 当以集群模式运行时，读取cluster的配置，以追加的方式写PID文件
func RecordPID() error {
	// 集群模式运行时 记录的是master group的PID
	// 停止master时 所有实例都会停止
	pid := os.Getpid()
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "logs", "jjgo.pid")
	logger.JJGoLogger.Info("JJGo running PID was written into ./logs/jjgo.pid")
	err := ioutil.WriteFile(file, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		logger.JJGoLogger.Error("JJGo Record PID failed", err)
		return err
	}
	return nil
}
