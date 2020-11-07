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

var WhiteList = model.WhiteList{}

// 校验白名单
func ReadWhite() {
	cwd, _ := os.Getwd()
	whitePath := path.Join(cwd, "conf", "white.json")
	f , err := ioutil.ReadFile(whitePath)
	if err != nil {
		return
	}
	_ = json.Unmarshal(f, &WhiteList)

}
