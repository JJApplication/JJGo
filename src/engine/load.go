/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

加载路由 中间件 配置
*/
package engine

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/api/html"
	"jjgo/src/api/jjgo"
	"jjgo/src/api/jjmail"
	"jjgo/src/api/mysite"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/logger"
	"jjgo/src/middleware"
)


// 加载路由
func loadRouter(handler *gin.Engine) {
	html.JJGoSwagger(handler)
	html.JJGoIndex(handler)
	jjgo.JJGoSelf(handler)
	jjmail.JJMail(handler)
	mysite.Mysite(handler)
}

// 加载中间件
func loadMiddleware(handler *gin.Engine) {
	handler.Use(middleware.JJGoLog())
	handler.Use(gin.Recovery())
	handler.Use(middleware.ResponseTime())
	handler.Use(middleware.APICount())
	handler.Use(middleware.Cors())
	handler.Use(middleware.Firewall())
	handler.Use(middleware.PageNotFound())
	// 分应用加载auth中间件
}

// 配置GIN
func GINConf() {
	logger.JJGoLogger = logger.InitLogger(config.JJGoConf.LogPath, config.JJGoConf.Mode)
	gin.SetMode(config.JJGoConf.Mode)
	gin.DisableConsoleColor()
}

// 配置中间件的黑白名单
func JJGoConf() {
	config.ReadWhite()
	config.ReadBlack()
	config.InitJJGoVersion()
	config.InitChangeLog()
}

// 加载数据库
// 因为使用多个数据库，所以默认的全局加载使用jjgo数据库
func LoadDB() jjgorm.JJGorm {
	return jjgorm.JJGormClient()
}