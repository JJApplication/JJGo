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

	if *arg == "nothing" {
		runMode := config.JJGoConf.RunMode
		if runMode == "standalone" || runMode == "single" {
			switch *run {
			case "default":
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
	}

	engine.ParseArgs(*arg)
}
