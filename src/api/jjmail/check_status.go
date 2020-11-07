/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package jjmail

import (
	"os/exec"
)

func checkStatus() bool {
	// 优先检查系统的python版本
	shell := "pgrep celery"
	cmd := exec.Command("bash", "-c", shell)
	opt, err := cmd.Output()
	if err != nil {
		return false
	}else {
		if string(opt[:]) != "" || len(opt) > 0 {
			return true
		}
	}
	return false
}