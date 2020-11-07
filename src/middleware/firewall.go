/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

拦截垃圾请求
支持白名单 黑名单
*/
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/model"
	"jjgo/src/util"
)

// 根据useragent refer host ip拦截请求
func Firewall() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 直接引入白名单和黑名单，因为在注册中间件之前已经初始化了名单
		blackList := config.BlackList
		whiteList := config.WhiteList

		refer := c.Request.Referer()
		host := c.Request.Host
		ip := c.ClientIP()
		// 先通过全部白名单的请求
		if contains(whiteList.Refer, refer) || contains(whiteList.Host, host) || contains(whiteList.IP, ip) {
			c.Next()
		}else {
			// 对黑名单直接拦截
			if contains(blackList.Refer, refer) || contains(blackList.Host, host) || contains(blackList.IP, ip) {
				util.JJResponse(
					c,
					model.HTTP_FORBIDDEN,
					// 因为想要输出错误页面信息
					"forbidden to access",
					"403 Forbidden",
				)
				c.Abort()
				return
			}
			c.Next()
		}
	}
}

func contains(sl []string, s string) bool {
	for _, v := range sl {
		if v == s || strings.Contains(s, v) || strings.Contains(v, s) {
			return true
		}
	}
	return false
}
