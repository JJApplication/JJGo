/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"jjgo/src/model"
)

var BlackList = model.BlackList{}

// 校验白名单
func ReadBlack() {
	cwd, _ := os.Getwd()
	whitePath := path.Join(cwd, "conf", "black.json")
	f , err := ioutil.ReadFile(whitePath)
	if err != nil {
		return
	}
	_ = json.Unmarshal(f, &BlackList)
}

