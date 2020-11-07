/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package cmd

import (
	"os/exec"

	"github.com/color"
)

func CliStatus() {
	shell1 := "pgrep jjgo"
	cmd := exec.Command("bash", "-c", shell1)
	opt, _ := cmd.Output()
	pid := string(opt[:])
	if pid == "" {
		color.Red("JJGo服务没有在运行")
	}

	color.Green("JJGo服务运行在PID: %s", pid)
}
