/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 统一响应函数
// rest接口
// 返回状态码， 响应信息， 数据， 响应时间（毫秒）
func JJResponse(c *gin.Context, httpCode int, msg string, data interface{}) {
	// 可能存在请求直接拦截的情况 时间差为空
	var latency int64
	if _, ifExit := c.Get("start");ifExit {
		latency = time.Since(c.GetTime("start")).Nanoseconds() /1e6
	}else {
		latency = 0
	}

	c.JSON(httpCode, gin.H{
		"code": httpCode,
		"time": fmt.Sprintf("%vms", latency),
		"msg": msg,
		"data": data,
	})
	return
}

// 统一raw格式 JSON响应
func JJRawResponse(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
	return
}