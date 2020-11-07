/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gookit/color"
)

var logPath string
var logFile  = "jjgo.log"

// JJGo日志分析程序
func main() {
	args := os.Args
	if len(args) <2 {
		color.Red.Println("需要至少一个参数")
		color.BgBlue.Println("输入[help] 查看使用方式")
		return
	}
	switch args[1] {
	case "help":
		color.Green.Println("JJLog Usage: >>>")
		color.Green.Println("[help] 帮助信息")
		color.Green.Println("[show] 查看日志")
		color.Green.Println("[clean] 清空日志")
		color.Green.Println("[search] 搜索日志")

	case "show":
		show()

	case "clean":
		clean()

	case "search":
		if len(args) >= 3 {
			search(args[2])
		}else {
			color.Red.Println("请指定search字符串")
		}
	}
}

func checkFile(path string) bool {
	if _, err :=os.Stat(path);err != nil {
		return false
	}else {
		return true
	}
}

func show() {
	cwd, _ := os.Getwd()
	logPath = path.Join(cwd, "logs", logFile)
	if !checkFile(logPath) {
		color.BgRed.Println("日志路径有无，jjlog应该运行在jjgo同级目录下")
		return
	}
	file, err := ioutil.ReadFile(logPath)
	if err != nil {
		color.BgRed.Println("日志无法读取")
	}else {
		color.Green.Printf("%s", file)
	}
}

func clean() {
	cwd, _ := os.Getwd()
	logPath = path.Join(cwd, "logs", logFile)
	if !checkFile(logPath) {
		color.BgRed.Println("日志路径有无，jjlog应该运行在jjgo同级目录下")
		return
	}
	err := ioutil.WriteFile(logPath, []byte(""), 0644)
	if err != nil {
		color.BgRed.Println("日志无法清空")
	}else {
		color.BgGreen.Println("日志已清空")
	}
}

func search(sub string) {
	cwd, _ := os.Getwd()
	logPath = path.Join(cwd, "logs", logFile)
	if !checkFile(logPath) {
		color.BgRed.Println("日志路径有无，jjlog应该运行在jjgo同级目录下")
		return
	}
	file, err := ioutil.ReadFile(logPath)
	if err != nil {
		color.BgRed.Println("日志无法读取")
	}else {
		content := string(file)
		count := strings.Count(content, sub)
		color.BgGreen.Printf("匹配次数: %d\n", count)
		contentList := strings.Split(content, "\n")
		for _, line := range contentList {
			if strings.Contains(line, sub) {
				color.BgGreen.Println(line)
			}
		}
	}
}