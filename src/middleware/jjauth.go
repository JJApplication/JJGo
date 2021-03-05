/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

认证中间件
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/util"
)

// Auth认证方式
// 支持token头header params验证
// 普通验证方式无预先申请token接口，输入自定的token key直接验证(明文校验)
//
// 注意验证是在成功登陆或者请求时的操作，而firewall是对垃圾请求的拦截所以应该在auth前注册
//
// jjauth jjapps之间的直接通信协议 在白名单前直接被验证
// 支持cookie token-header param三种方式
// 实现任意一种即可其中cookie内容需要加解密 仅在服务器配置中开启
func JJAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jjauth校验失败则进入顺序其他校验流程
		expectedKey := config.JJGoConf.JJAuthKey
		if util.VerifyAgent(c) && util.VerifyAll(c, expectedKey) {
			c.Header("jjauth", "active")
			c.Next()
		}else {
			// first firewall
			// CheckList(c) 多余校验firewall作为顶层足够
			// second token auth
		}
	}
}