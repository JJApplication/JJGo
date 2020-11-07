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

func CliStop() {
	shell1 := "cat ../logs/jjgo.pid"
	cmd := exec.Command("bash", "-c", shell1)
	opt, _ := cmd.Output()
	pid := string(opt[:])
	shell2 := "kill -2 " + pid

	cmd = exec.Command("bash", "-c", shell2)
	err := cmd.Run()
	if err != nil {
		color.Red("停止JJGo服务失败, %s", err)
	}else {
		color.Cyan("停止JJGo服务成功")
	}
}
