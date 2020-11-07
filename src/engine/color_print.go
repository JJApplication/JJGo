/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

定义颜色输出样式
*/
package engine

import (
	"fmt"

	"github.com/gookit/color"
)

// 仅可包内使用
//
// 前景色 背景色
// 带背景色的部分 最后需要解除背景色
type console struct {
	Version string
}

func (*console) FgRed(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.Red.Println(format)
	}else {
		color.Red.Println(fmt.Sprintf(format, a))
	}
}

func (*console) FgGreen(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.Green.Println(format)
	}else {
		color.Green.Println(fmt.Sprintf(format, a))
	}
}

func (*console) FgCyan(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.Cyan.Println(format)
	}else {
		color.Cyan.Println(fmt.Sprintf(format, a))
	}
}

func (*console) FgYellow(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.Yellow.Println(format)
	}else {
		color.Yellow.Println(fmt.Sprintf(format, a))
	}
}

func (*console) BgRed(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.BgRed.Println(format)
	}else {
		color.BgRed.Println(fmt.Sprintf(format, a))
	}
}

func (*console) BgGreen(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.BgGreen.Println(format)
	}else {
		color.BgGreen.Println(fmt.Sprintf(format, a))
	}
}

func (*console) BgCyan(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.BgCyan.Println(format)
	}else {
		color.BgCyan.Println(fmt.Sprintf(format, a))
	}
}

func (*console) BgYellow(format string, a ...interface{}) {
	if len(a) <= 0 {
		color.BgYellow.Println(format)
	}else {
		color.BgYellow.Println(fmt.Sprintf(format, a))
	}
}
