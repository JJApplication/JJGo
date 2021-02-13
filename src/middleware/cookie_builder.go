/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/util"
)

// cookie生成
// 作用为jjgo统计来自一个用户的接口请求信息
// 需要使用cookie访问的限制性内容

func CookieBuilder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 内建的cookie主要作用统计用户的唯一性
		// 注意：用于app间交互的cookie是有client端携带的不由jjgo生成
		cook := util.CookieNew(c)
		c.SetCookie("userid", cook, 3600, "/", "localhost", false, true)
		c.Next()
	}
}