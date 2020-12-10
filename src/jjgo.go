/*
App: JJGo
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-20
*/
package main

import (
	"flag"
	"jjgo/src/engine"
)

func main() {
	runFlag  := flag.String("r", "default", "默认模式|全新模式|守护模式｜测试模式")
	sFlag  := flag.String("s", "", "[start][stop][restart][reload][status]")
	editFlag := flag.Bool("e", false, "修改jjgo配置文件")
	cleanFlag := flag.Bool("c", false, "清空jjgo日志")
	showFlag := flag.Bool("l", false, "打印jjgo日志 指定行数")
	versionFlag := flag.Bool("v", false, "显示jjgo版本")
	helpFlag := flag.Bool("h", false, "显示帮助信息")

	flag.Parse()
	// 避免出现jjgo的多实例运行 默认参数时无操作
	if *editFlag {
		engine.EditConf()
		return
	}
	if *cleanFlag {
		engine.ClearLog()
		return
	}
	if *showFlag {
		// 此时应该有args
		line := flag.Arg(0)
		engine.ShowLog(line)
		return
	}
	if *versionFlag {
		engine.GetVersion()
		return
	}
	if *helpFlag {
		flag.Usage()
		return
	}
	if *sFlag == "" {
		engine.ServiceJJGo(*runFlag)
	}else {
		engine.ParseArgs(*sFlag)
	}
}
