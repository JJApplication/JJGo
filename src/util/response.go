/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 统一响应函数
// rest接口
// 返回状态码， 响应信息， 数据， 响应时间（毫秒）
func JJResponse(c *gin.Context, httpCode int, msg string, data interface{}) {
	latency := time.Since(c.GetTime("start")).Nanoseconds() /1e6
	c.JSON(httpCode, gin.H{
		"code": httpCode,
		"time": fmt.Sprintf("%vms", latency),
		"msg": msg,
		"data": data,
	})
}

// 统一raw格式 JSON响应
func JJRawResponse(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
}