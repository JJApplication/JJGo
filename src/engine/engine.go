/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

生成web Engine
*/
package engine

import (
	"context"
	"fmt"
	"jjgo/src/model/errors"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/logger"
	"jjgo/src/util"
)

var handler *gin.Engine
var JJGorm jjgorm.JJGorm

type JJGoEngine struct {
	engine *gin.Engine
	pid int
	// cluster props
	ID  int
	port string
}

// 加载配置
// 加载env


// return a jjgo engine
//
// GIN基础配置由ENV环境变量加载
// jjgo基础配置通过ini文件加载
// 避免循环引入在开始服务前新建日志实例

func JJGo() JJGoEngine {
	var jjgo JJGoEngine
	// 注册在路由前 保证路由里可以使用数据库
	JJGorm = LoadDB()
	JJGorm.Connect(config.JJGoConf.DBJJGo)
	// 在停止时保证数据库正常关闭
	GINConf()
	JJGoConf()
	handler = gin.New()
	loadMiddleware(handler)
	loadRouter(handler)

	jjgo.engine = handler
	jjgo.pid = os.Getpid()

	return jjgo
}

// http server运行jjgo
func (jjgo *JJGoEngine) Run() {
	con.FgGreen("以正常模式启动JJGo服务")
	logger.JJGoLogger.Info("以正常模式运行...")
	s := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d",config.JJGoConf.Port),
		Handler: jjgo.engine,
		ReadTimeout: config.JJGoConf.ReadTimeout,
		WriteTimeout: config.JJGoConf.WriteTimeout,
		IdleTimeout: config.JJGoConf.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = util.RecordPID()
	var sigChan chan os.Signal
	sigChan = make(chan os.Signal)
	go RegisterSignal(s, sigChan)
	err := s.ListenAndServe()
	if err != nil {
		con.FgRed("以正常模式启动JJGo服务失败")
		logger.JJGoLogger.Error(errors.SERVER_START_FAIL, err)
	}
	<-sigChan
}

// 运行测试模式8080
func (jjgo *JJGoEngine) RunTest() {
	con.FgGreen("以测试模式启动JJGo服务")
	logger.JJGoLogger.Info("以测试模式运行...")
	//err := jjgo.engine.Run(fmt.Sprintf(":%d", config.JJGoConf.Port))
	gin.SetMode("debug")
	err := jjgo.engine.Run(":8080")
	if err != nil {
		con.FgRed("以测试模式启动JJGo服务失败")
		logger.JJGoLogger.Error(errors.SERVER_START_FAIL, err)
		os.Exit(1)
	}
}

// 优雅的守护进程
func (jjgo *JJGoEngine) RunDaemon() {
	runTimePID := syscall.Getppid()
	if runTimePID == 1 {
		// 是守护进程
		return
	}
	con.FgGreen("以守护模式启动JJGo服务")
	logger.JJGoLogger.Info("以守护进程启动")
	// 守护进程额外写入daemon.txt文件
	fp, err := os.OpenFile(path.Join("logs", "daemon.txt"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		con.FgRed("以守护模式启动JJGo服务失败")
		logger.JJGoLogger.Error("守护进程启动失败", err)
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()

	fp.WriteString("以守护模式启动，运行PID保存至./logs/jjgo.pid\n")
	// 这里参数只应该传arg[0] 否则会进入无限循环开启守护进程
	cmd := exec.Command(os.Args[0])
	cmd.Stdout = fp
	cmd.Stderr = fp
	cmd.Stdin = nil
	// only work on linux
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	if err := cmd.Start();err != nil {
		con.FgRed("以守护模式启动JJGo服务失败")
		logger.JJGoLogger.Error("守护进程exec启动失败", err)
		panic(err)
	}
	os.Exit(0)
}

// 清空JJGo日志 获取一个全新的Engine
func (jjgo *JJGoEngine) CleanRun() {
	con.FgGreen("以全新模式启动JJGo服务")
	logger.JJGoLogger.Info("以全新模式运行...")
	logger.JJGoLogger.Clean()
	s := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d",10086),
		Handler: jjgo.engine,
		ReadTimeout: config.JJGoConf.ReadTimeout,
		WriteTimeout: config.JJGoConf.WriteTimeout,
		IdleTimeout: config.JJGoConf.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = util.RecordPID()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			con.FgRed("以全新模式启动JJGo服务失败")
			logger.JJGoLogger.Error(errors.SERVER_START_FAIL, err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.JJGoLogger.Warning(errors.SERVER_STOPPED, err.Error())
		os.Exit(0)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.JJGoLogger.Info(errors.SERVER_EXIT)
	}
}