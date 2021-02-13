/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package util

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// cookie的内建方法 用于cookie的key-value获取

// cookie的生成
// 根据用户的IP地址加uuid
func CookieNew(c *gin.Context) string {
	// first parse
	inParseCookie := CookieParse(c)
	if inParseCookie != "" {
		return inParseCookie
	}
	ip := c.ClientIP()
	if ip == "" {
		ip = "127.0.0.1"
	}
	uuidNew, err := uuid.NewV4()
	if err != nil {
		uuidNew = uuid.Nil
	}
	keyRaw := ip + uuidNew.String()
	return keyRaw
}

// parse cookie in all cookies
func CookieParse(c *gin.Context) string {
	cookieUser, non := c.Cookie("userid")
	if non != nil {
		cookieUser = ""
	}
	return cookieUser
}
