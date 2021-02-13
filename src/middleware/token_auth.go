/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/model"
	"jjgo/src/util"
)

// 常规的token验证
// 支持token头header params验证
// 普通验证方式无预先申请token接口，输入自定的token key直接验证(明文校验)
//
// 注意验证是在成功登陆或者请求时的操作，而firewall是对垃圾请求的拦截所以应该在auth前注册
//
// 作为部分接口的引入中间件 只需要在此中间件前添加jjauth即可
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedKey := config.JJGoConf.JJAuthKey
		if util.VerifyAgent(c) && util.VerifyAll(c,expectedKey) {
			c.Header("jjauth", "active")
			c.Next()
		}else {
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
}

func NormalTokenAuth(c *gin.Context) {
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
