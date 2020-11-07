/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

响应时间中间件
*/
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 计算响应时间差
//
// 所有响应均使用到了时间差 所以此中间件优先级高
func ResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 向上下文写入参数
		start := time.Now()
		c.Set("start", start)
		c.Next()
	}
}
