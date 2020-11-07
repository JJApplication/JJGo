/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

接口次数统计
*/
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/model/database"
)

var LastTime int64
var db jjgorm.JJGorm
var count int

func init() {
	db = jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBJJGo)
	// 初始化时间
	LastTime = time.Now().Unix()
	count = 0
}
func APICount() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 每次读写数据库占用资源
		// 每5min保存一次
		// 需要新建一个数据库 避免循环引入
		// new := database.JJGoAPICount{}
		// db.CreateTable(&new)

		if time.Now().Unix() - LastTime <= 1 * 60 {
			// 小于差值 先相加
			count++
		}else {
			n := database.JJGoAPICount{}
			err := db.FindFirstByStruct(database.JJGoAPICount{API: "all"}, &n).Error
			if err == nil {
				n.Count += count
				count = 0 // 重新赋值
				LastTime = time.Now().Unix()
				db.Update(database.JJGoAPICount{}, &n)
			}else {
				db.Insert(&database.JJGoAPICount{API: "all", Count: 1})
			}
		}

		c.Next()
	}
}

