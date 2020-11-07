/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

集群模式运行
*/
package engine

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
	"jjgo/src/config"
	"jjgo/src/logger"
	"jjgo/src/util"
)

var JJGoCluster []JJGoEngine
var group errgroup.Group

// 根据配置文件定义的cluster自动配置
func Cluster() {
	clusterConfig := config.JJGoConf.Cluster
	if len(clusterConfig) <= 0 {
		logger.JJGoLogger.Error("以集群模式启动失败，将以独立模式运行", nil)
		var jjgoEngine JJGoEngine
		jjgoEngine = JJGo()
		jjgoEngine.Run()
	}else {
		// 读取cluster的port 确定需要启动的服务数
		for id, port := range clusterConfig {
			engine := JJGo()
			engine.ID = id
			engine.port = port
			JJGoCluster = append(JJGoCluster, engine)
		}
		// 异步启动
		for _, engine := range JJGoCluster {
			s := StartCluster(engine)
			group.Go(func() error {
				return s.ListenAndServe()
			})
		}

		if err := group.Wait();err != nil {
			logger.JJGoLogger.Error("集群运行错误%s", err)
		}
	}
}

// 添加server组
func StartCluster(eng JJGoEngine) *http.Server {
	logger.JJGoLogger.Info(fmt.Sprintf("以集群模式运行, 实例ID: %d, 端口: %s", eng.ID, eng.port))
	s := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%s",eng.port),
		Handler: eng.engine,
		ReadTimeout: config.JJGoConf.ReadTimeout,
		WriteTimeout: config.JJGoConf.WriteTimeout,
		IdleTimeout: config.JJGoConf.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = util.RecordPID()

	return s
}