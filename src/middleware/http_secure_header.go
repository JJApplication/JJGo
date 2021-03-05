/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"github.com/gin-gonic/gin"
)

// http安全响应头
func HTTPSafeHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("JJGO Server", "JJGO-API")
		c.Header("Version", "API V3")
		c.Header("APP", "JJGO")
		c.Header("Copyright", "Landers1037")
		c.Header("Github", "github.com/landers1037")
		c.Header("X-frame-Options", "SAMEORIGIN")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Cache-Control", "no-cache")
		c.Next()
	}
}
