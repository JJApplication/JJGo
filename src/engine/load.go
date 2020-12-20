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
	// 读取配置文件字典 动态加载中间件 只有全局的可以生效
	// 中间件加载字典
	middlewareMap := map[string] gin.HandlerFunc {
		"log": middleware.JJGoLog(),
		"recovery": gin.Recovery(),
		"type_allowed": middleware.TypeAllowed(),
		"response_time": middleware.ResponseTime(),
		"api_count": middleware.APICount(),
		"cors": middleware.Cors(),
		"page_not_found": middleware.PageNotFound(),
	}

	if config.JJGoConf.MiddleWare == nil || len(config.JJGoConf.MiddleWare) == 0 {
		loadAll(handler)
	}else {
		// 对字典重排 保证中间件加载顺序
		middle := sortMiddleware(config.JJGoConf.MiddleWare)
		for _, m := range middle {
			// start load
			middleKey := m["middle"]
			val := m["val"]
			mid, ok := middlewareMap[middleKey]
			// 此中间件存在
			if ok && val != "0" && val != "false" && val != "" {
				handler.Use(mid)
			}
		}
		// 分应用加载auth中间件
	}
}

func sortMiddleware(m map[string]string) []map[string]string {
	// 默认的map不排序 我们自己设置带顺序的数组
	sortedList := []string{
		"log", "recovery", "type_allowed", "response_time", "api_count", "cors", "page_not_found", "jjauth"}
	var sorted []map[string]string

	for _, key := range sortedList {
		if _, ok := m[key];ok {
			sorted = append(sorted, map[string]string{"middle": key, "val": m[key]})
		}else {
			sorted = append(sorted, map[string]string{"middle": key, "val": ""})
		}
	}

	return sorted
}

// 默认加载全部
func loadAll(handler *gin.Engine) {
	handler.Use(middleware.JJGoLog())
	handler.Use(gin.Recovery())
	handler.Use(middleware.TypeAllowed())
	handler.Use(middleware.ResponseTime())
	handler.Use(middleware.APICount())
	handler.Use(middleware.Cors())
	handler.Use(middleware.Firewall())
	handler.Use(middleware.PageNotFound())
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