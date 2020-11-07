/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package engine

import (
	"fmt"
	"os/exec"

	"jjgo/src/config"
)

// 解析外部args参数 发送信号量
var con = console{Version: "1.0"}

func ParseArgs(arg string) {
	con.FgCyan("你正在使用JJGo内置的CLI命令")
	switch arg {
	case "stop":
		con.FgCyan("即将停止JJGo服务...")
		stop()
	case "restart":
		con.FgCyan("即将重启JJGo服务...")
	case "start":
		con.FgCyan("JJGo服务启动中...")
		start()
	case "reload":
		con.FgCyan("JJGo服务重新加载中...")
		reload()
	case "help":
		help()
	default:
		con.FgGreen("Usage: '-s' [start][stop][restart][reload]")
		con.FgGreen("Type -s help for more...")
		con.FgRed("Invalid Args...")
	}
}

func start() {
	con.BgCyan("仅支持默认启动，守护进程请使用")
	con.BgGreen("jjgo -r daemon")
	cmd := exec.Command("bash", "-c", "nohup ./jjgo > /dev/null 2>&1 &")
	err := cmd.Run()
	if err != nil {
		con.BgRed("JJGo启动失败, %v", err)
		return
	}
	// 关闭启动的协程
	// 不是用脚本启动的没有协程
	con.BgGreen("JJGo启动完毕")
}

func stop() {
	pidPath := config.JJGoConf.LogRoot
	pidr, err := exec.Command("bash", "-c", fmt.Sprintf("cat %s/jjgo.pid", pidPath)).Output()
	if err != nil {
		con.BgRed("读取JJGo PID文件失败, %s", pidPath)
		con.BgYellow("错误提示: %v", err)
		return
	}
	pid := string(pidr[:])
	opt , err := exec.Command("bash", "-c", fmt.Sprintf("kill -2 %s", pid)).Output()
	if err != nil {
		con.BgRed("发送SIGINT信号失败, PID: %s", pid)
		con.BgYellow("错误提示: %v", err)
		return
	}
	if len(opt) <= 0 {
		con.BgGreen("JJGo服务已停止")
	}else {
		con.BgYellow("JJGo服务停止失败，请检查日志")
	}
}

func reload() {
	pidPath := config.JJGoConf.LogRoot
	pidr, err := exec.Command("bash", "-c", fmt.Sprintf("cat %s/jjgo.pid", pidPath)).Output()
	if err != nil {
		con.BgRed("读取JJGo PID文件失败, %s", pidPath)
		con.BgYellow("错误提示: %v", err)
		return
	}
	pid := string(pidr[:])
	// USR1 10 USR2 12
	opt , err := exec.Command("bash", "-c", fmt.Sprintf("kill -10 %s", pid)).Output()
	if err != nil {
		con.BgRed("发送SIGUSR信号失败, PID: %s", pid)
		con.BgYellow("错误提示: %v", err)
		return
	}
	if len(opt) <= 0 {
		con.BgGreen("JJGo服务重载完成")
	}else {
		con.BgYellow("JJGo服务重载失败，请检查日志")
	}
}

func help() {
	con.BgCyan("JJGo -s HELP")
	con.BgGreen("[start] 默认模式启动jjgo ")
	con.BgGreen("[stop] 停止jjgo进程")
	con.BgGreen("[restart] 重启jjgo服务")
	con.BgGreen("[reload] 重新加载jjgo配置")
}