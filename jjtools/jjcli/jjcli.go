/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package main

import (
	"os"

	"github.com/color"
	"jjgo/jjtools/jjcli/cmd"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		color.Yellow("JJCLI - CLI for JJGo\nUsage: [run|cleanrun|stop|status]\n" +
			"- run 标准模式运行JJGo服务\n" +
			"- cleanrun 全新模式运行JJGo服务\n" +
			"- stop 停止JJGo服务\n" +
			"- status 查看JJGo服务状态")
	}else {
		switch args[1] {
		case "run":
			cmd.CliRun()
		case "cleanrun":
			cmd.CliCleanRun()
		case "stop":
			cmd.CliStop()
		case "status":
			cmd.CliStatus()
		}
	}
}

