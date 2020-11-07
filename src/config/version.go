/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package config

import (
	"jjgo/src/model"
	"jjgo/src/util"
)

var JJGoVersion model.JJGoVersion

func InitJJGoVersion() {
	JJGoVersion = util.ReadVersion()
}
