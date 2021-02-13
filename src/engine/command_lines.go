/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package engine

import (
	"jjgo/src/config"
	"jjgo/src/logger"
	"jjgo/src/util"
	"os"
	"os/exec"
	"path"
)

// 命令行驱动
func EditConf() {
	cwd, _ := os.Getwd()
	confPath := path.Join(cwd, "conf", "jjgo.ini")
	_, err := os.Stat(confPath)
	if err != nil {
		con.FgRed("打开配置文件失败")
		logger.JJGoLogger.Error("打开配置文件失败, err:%s\n", err.Error())
		return
	}
	cmd := exec.Command("vi", confPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	return
}

func ClearLog() {
	cwd, _ := os.Getwd()
	logPath := path.Join(cwd, "logs", "jjgo.log")
	_, err := os.Stat(logPath)
	if err != nil {
		con.FgRed("打开日志文件失败")
		logger.JJGoLogger.Error("打开日志文件失败, err:%s\n", err.Error())
		return
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		con.FgRed("日志文件清空失败, err:%s\n", err.Error())
		return
	}
	_, _ = f.WriteString("")
	_ = f.Close()
	con.FgGreen("日志已经被清空")
	return
}

func ShowLog(line string) {
	cwd, _ := os.Getwd()
	logPath := path.Join(cwd, "logs", "jjgo.log")
	_, err := os.Stat(logPath)
	if err != nil {
		con.FgRed("打开日志文件失败")
		logger.JJGoLogger.Error("打开日志文件失败, err:%s\n", err.Error())
		return
	}
	if line == "" {
		cmd := exec.Command("cat", logPath)
		res, _ := cmd.Output()
		con.FgGreen(string(res))
	}else {
		cmd := exec.Command("tail", "-n", line, logPath)
		res, _ := cmd.Output()
		con.FgGreen(string(res))
	}
	return
}

func GetVersion() {
	jjgoVersion := util.ReadVersion()
	con.FgGreen("Name: JJGo\nVersion: %s\nBuild: %s\nAPIServer: %s\n",
		jjgoVersion.Version, jjgoVersion.BuildDate, jjgoVersion.APIServer)
	//data, _ := json.MarshalIndent(jjgoVersion, "", "\t")
	//fmt.Printf("%s", string(data))
	return
}

func ServiceJJGo(runFlag string) {
	con.FgCyan("请指定启动参数-s | -r, 查看帮助-h")
	runMode := config.JJGoConf.RunMode
	if runMode == "standalone" || runMode == "single" {
		switch runFlag {
		case "default":
			// 改为无操作 避免多实例
			jjgoEngine := JJGo()
			jjgoEngine.Run()
		case "clean":
			jjgoEngine := JJGo()
			jjgoEngine.CleanRun()
		case "daemon":
			jjgoEngine := JJGo()
			jjgoEngine.RunDaemon()
		case "test":
			jjgoEngine := JJGo()
			jjgoEngine.RunTest()
		default:
			con.FgRed("输入的参数无效，请使用-h获取帮助信息")
			return
		}
	}else {
		// using cluster
		con.FgYellow(config.JJGoConf.RunMode)
		Cluster()
	}
}