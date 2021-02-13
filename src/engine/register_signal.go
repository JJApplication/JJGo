/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

注册信号量
*/
package engine

import (
	"context"
	"jjgo/src/middleware"
	"jjgo/src/model/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"jjgo/src/config"
	"jjgo/src/logger"
)

func RegisterSignal(s *http.Server, sigChan chan os.Signal) {
	// 监听信号量
	// ONLY WORK ON LINUX
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)
	// select判断
	// kill (no param) default send syscall.SIGTERM
	// kill -1 SIGHUP
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	// kill -10 -12 SIGUSR1 SIGUSR2
	for sig := range sigChan {
		switch sig {
		case syscall.SIGINT :
			// 将数据库数据保存
			middleware.ForceSaveCount()
			// 保证数据库释放
			logger.JJGoLogger.Info("数据库句柄已关闭")
			JJGorm.Close()

			// 停止
			if err := s.Shutdown(context.Background()); err != nil {
				con.FgGreen("JJGo服务关闭完成")
				logger.JJGoLogger.Warning(errors.SERVER_STOPPED, err.Error())
				logger.JJGoLogger.Info(errors.SERVER_EXIT)
			}
			// catching ctx.Done(). timeout of 5 seconds.
			logger.JJGoLogger.Info(errors.SERVER_EXIT)
			logger.JJGoLogger.Warning(errors.SERVER_STOPPED)
			os.Exit(0)

		case syscall.SIGTERM:
			// 强制关闭
			middleware.ForceSaveCount()
			con.FgGreen("JJGo服务强制关闭完成")
			logger.JJGoLogger.Info(errors.SERVER_TERMINATED)

		case syscall.SIGUSR1:
			// 先引入所有的json文件
			// 当前的json配置只有白名单和黑名单, 更新日志, 版本号
			// 新增对配置文件中的中间件的重新加载
			con.FgGreen("JJGo服务重载中...")
			config.ReadWhite()
			config.ReadBlack()
			config.InitChangeLog()
			config.InitJJGoVersion()
			config.ReloadMiddleConf()
			con.FgGreen("JJGo服务重载完成")
			logger.JJGoLogger.Info(errors.SERVER_RELOADED)
		}
	}
}
