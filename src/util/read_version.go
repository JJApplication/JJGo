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

// 读取JJGo Version
func ReadVersion() model.JJGoVersion {
	cwd, _ := os.Getwd()
	version := path.Join(cwd, "conf", "version.json")
	f, err := ioutil.ReadFile(version)
	if err != nil {
		logger.JJGoLogger.Error("更新日志文件读取失败", err)
		return model.JJGoVersion{}
	}
	var res model.JJGoVersion
	_ = json.Unmarshal(f, &res)

	return res
}
