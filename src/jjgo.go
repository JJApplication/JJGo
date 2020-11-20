/*
App: JJGo
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-20
*/
package main

import (
	"flag"
	"fmt"

	"jjgo/src/config"
	"jjgo/src/engine"
)

func main() {
	run  := flag.String("r", "default", "默认模式|全新模式|守护模式")
	arg  := flag.String("s", "nothing", "[start][stop][restart][reload]")
	flag.Parse()
	// 避免出现jjgo的多实例运行 默认参数时无操作
	if *arg == "nothing" {
		runMode := config.JJGoConf.RunMode
		if runMode == "standalone" || runMode == "single" {
			switch *run {
			case "default":
				// 改为无操作 避免多实例
				fmt.Println("请指定启动参数-s | -r, 查看帮助")
				fmt.Println("输入jjgo -s help以查看帮助")
				jjgoEngine := engine.JJGo()
				jjgoEngine.Run()
			case "clean":
				jjgoEngine := engine.JJGo()
				jjgoEngine.CleanRun()
			case "daemon":
				jjgoEngine := engine.JJGo()
				jjgoEngine.RunDaemon()
			default:
				jjgoEngine := engine.JJGo()
				jjgoEngine.Run()
			}
		}else {
			// using cluster
			fmt.Println(config.JJGoConf.RunMode)
			engine.Cluster()
		}
	}else {
		engine.ParseArgs(*arg)
	}
}
