/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

JJGo的self接口，主要负责和前端交互json数据
*/
package jjgo

import (
	"jjgo/src/url"
	"os"

	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/model"
	"jjgo/src/model/database"
	"jjgo/src/model/entity"
	"jjgo/src/util"
)

var ApiJJGo *gin.RouterGroup

func JJGoSelf(r *gin.Engine) {
	info := func(c *gin.Context) {
		util.JJResponse(
			c,
			model.HTTP_STATUS_OK,
			"this's API for JJGo",
			map[string]string{"api": "[jjgo]", "path": c.FullPath()},
		)
	}
	ApiJJGo = r.Group(url.PREFIX_JJGO)
	{
		ApiJJGo.GET(url.JJGO_INDEX, info)
		ApiJJGo.GET(url.JJGO_PUB, jjgoPubTest)
		ApiJJGo.GET(url.JJGO_STATUS, jjgoStatus)
		ApiJJGo.GET(url.JJGO_CHANGELOG, jjgoChangeLog)
		ApiJJGo.GET(url.JJGO_VERSION, jjgoVersion)
		ApiJJGo.GET(url.JJGO_DEMO, jjgoDemo)
		ApiJJGo.GET(url.JJGO_SWAGGER, jjgoJSON)
	}
}

// 公共测试接口
func jjgoPubTest(c *gin.Context) {
	util.JJResponse(c, model.HTTP_STATUS_OK, "Test for JJGo public",
		entity.JJGoPubTest{
			Client_IP:    c.ClientIP(),
			Method:       c.Request.Method,
			Host:         c.Request.Host,
			Content_Type: "JSON",
		})
}

// 服务运行状态
// 获取状态和请求数量 从数据库读取
func jjgoStatus(c *gin.Context) {
	var count database.JJGoAPICount
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBJJGo)
	err := db.FindFirstByStruct(database.JJGoAPICount{API: "all"}, &count).Error
	defer db.Close()

	if err != nil {
		count = database.JJGoAPICount{}
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "jjgo status",
		entity.JJGoStatus{
			PID: os.Getpid(),
			Port: config.JJGoConf.Port,
			Count: count.Count,
		})
}

// 更新日志
func jjgoChangeLog(c *gin.Context) {
	util.JJResponse(c, model.HTTP_STATUS_OK, "jjgo change log",
		config.ChangeLog)
}

// 版本号
func jjgoVersion(c *gin.Context) {
	util.JJResponse(c, model.HTTP_STATUS_OK, "jjgo Version",
		config.JJGoVersion)
}

// api描述信息
func jjgoDemo(c *gin.Context) {
	util.JJResponse(c, model.HTTP_STATUS_OK, "jjgo demo",
		entity.JJGoDemo{
			Rest: map[string]string{"path": "/rest/*", "des": "restful接口"},
			File: map[string]string{"path": "/swagger/*", "des": "静态文件接口"},
			Html: map[string]string{"path": "/*.html", "des": "React前端页面"},
		})
}

// api json文档
func jjgoJSON(c *gin.Context) {
	// 因为调用接口响应为纯json 所以不能使用jjresponse
	data := util.SwaggerJson()
	util.JJRawResponse(c, model.HTTP_STATUS_OK, data)
}
