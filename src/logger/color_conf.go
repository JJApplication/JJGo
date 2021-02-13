/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package logger

import (
	"github.com/go-ini/ini"
	"os"
	"path"
)

// color config
func ColorConfig() bool {
	cwdPath , pathErr := os.Getwd()
	if pathErr != nil {
		return false
	}
	iniPath := path.Join(cwdPath, "conf" ,"jjgo.ini")

	cfg, _ := ini.Load(iniPath)
	color := cfg.Section("log").Key("color").MustBool(false)
	return color
}
