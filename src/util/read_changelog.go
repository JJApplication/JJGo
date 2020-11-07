/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"jjgo/src/logger"
	"jjgo/src/model"
)

// 读取本地的更新日志
func ReadChangeLog() []model.ChangeLog {
	cwd, _ := os.Getwd()
	changeLog := path.Join(cwd, "conf", "changelog.json")
	f, err := ioutil.ReadFile(changeLog)
	if err != nil {
		logger.JJGoLogger.Error("更新日志文件读取失败", err)
		return []model.ChangeLog{}
	}
	var res []model.ChangeLog
	_ = json.Unmarshal(f, &res)
	return res
}