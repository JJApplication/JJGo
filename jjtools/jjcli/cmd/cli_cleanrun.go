/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/color"
)

func CliCleanRun() {
	// 启动服务
	shell1 := "nohup ./jjgo -run clean > /dev/null 2>&1 &"
	cmd := exec.Command("bash", "-c", shell1)
	err := cmd.Run()

	shell2 := fmt.Sprintf("ps ax | grep \"%s\" | grep -v bash | grep -v grep | awk '{print $1}'", shell1)
	cmd = exec.Command("bash", "-c", shell2)
	opt, _ := cmd.Output()
	pid := string(opt[:])

	err = exec.Command("bash", "-c", "kill", pid).Run()
	if err != nil {
		color.Red("启动JJGo服务失败, %s", err)
	}else {
		color.Cyan("启动JJGo服务成功")
	}
}

