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
	"jjgo/src/model"
	"jjgo/src/util"
)

// Auth认证方式
// 支持token头header params验证
// 普通验证方式无预先申请token接口，输入自定的token key直接验证(明文校验)
//
// 注意验证是在成功登陆或者请求时的操作，而firewall是对垃圾请求的拦截所以应该在auth前注册
func JJAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMethod := config.JJGoConf.AuthMethod
		authKey := config.JJGoConf.AuthKey

		switch authMethod {
		case "token_header":
			if checkTokenHeader(c, authKey) {
				c.Next()
			}else {
				util.JJResponse(
					c,
					model.HTTP_FORBIDDEN,
					// 因为想要输出错误页面信息
					"forbidden to access, Token required",
					"403 Forbidden",
				)
				c.AbortWithStatus(model.HTTP_FORBIDDEN)
				return
			}
		case "token_params":
			if checkTokenParams(c, authKey) {
				c.Next()
			}else {
				util.JJResponse(
					c,
					model.HTTP_FORBIDDEN,
					// 因为想要输出错误页面信息
					"forbidden to access, Token required",
					"403 Forbidden",
				)
				c.AbortWithStatus(model.HTTP_FORBIDDEN)
				return
			}
		default:
			if checkTokenHeader(c, authKey) {
				c.Next()
			}else {
				util.JJResponse(
					c,
					model.HTTP_FORBIDDEN,
					// 因为想要输出错误页面信息
					"forbidden to access, Token required",
					"403 Forbidden",
				)
				c.AbortWithStatus(model.HTTP_FORBIDDEN)
				return
			}
		}
	}
}

func checkTokenHeader(c *gin.Context, authKey string) bool {
	headerToken := c.Request.Header.Get("token")
	if headerToken == util.TokenEncrypt(authKey) {
		return true
	}
	return false
}

func checkTokenParams(c *gin.Context, authKey string) bool {
	headerToken := c.DefaultQuery("token", "")
	if headerToken == util.TokenEncrypt(authKey) {
		return true
	}
	return false
}