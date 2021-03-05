/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"github.com/gin-gonic/gin"
)

// 开放api接口标准中间件
// 使用来jjtoken的用户通过邮箱name 获取到唯一token
// 根据域名判断所有权
// jjtoken分权机制 支持选择接口暴露
//
// 基于gin的中间件机制 该接口作为接口级中间件 优先作用于接口上
func JJTokenWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		jjtokenId := c.Request.Header.Get("jjtoken_id")
		jjtokenRaw := c.Request.Header.Get("jjtoken")

		if jjtokenId == "" || jjtokenRaw == "" {
			c.Next()
		}
		if jjtokenId == "superjj" {
			c.Next()
		}
	}
}
