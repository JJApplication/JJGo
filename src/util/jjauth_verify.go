/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// jjauth内部协议的验证函数
// cookie需要单独的配置支持，这里如果直接校验将会造成更大延时
// 所以cookie设计成双端协议 需要额外提供参数
// 所有的请求头 都应为jjclient字段

func VerifyAgent(c *gin.Context) bool {
	userAgent := c.Request.UserAgent()
	if strings.ToLower(userAgent) == "jjclient" {
		return true
	}
	return false
}

// cookie
// need if_cookie
func VerifyCookie(c *gin.Context) bool {
	if_cookie, err := c.Cookie("if_cookie")
	fmt.Println(if_cookie, err)
	return false
}

// auth token
// token name: jjauth
func VerifyToken(c *gin.Context, expectedKey string) bool {
	token := c.Request.Header.Get("jjauth")
	if token == "" {
		return false
	}else {
		// 支持明文和加密 建议加密
		if token == expectedKey {
			return true
		}
	}
	return false
}

// params
// params name: jjauth
func VerifyParams(c *gin.Context, expectedKey string) bool {
	param := c.Query("jjauth")
	if param == "" {
		return false
	}else {
		if param == expectedKey {
			return true
		}
	}
	return false
}

// all verify
func VerifyAll(c *gin.Context, expectedKey string) bool {
	return VerifyAgent(c) && VerifyCookie(c) || VerifyToken(c, expectedKey) || VerifyParams(c, expectedKey)
}