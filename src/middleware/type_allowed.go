/*
https type协议校验中间件
 */


package middleware

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/model"
	"jjgo/src/util"
	"strings"
)

// 仅支持json rest的请求
// 对get请求放行

func TypeAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctype := c.ContentType()
		method := c.Request.Method
		if method == "GET" || method == "get" {
			c.Next()
		}else {
			if strings.ToLower(ctype) == "application/json" {
				c.Next()
			}else {
				util.JJResponse(
					c,
					model.HTTP_FORBIDDEN,
					// 因为想要输出错误页面信息
					"only JSON type request can be accept",
					"403 Forbidden",
				)
				c.AbortWithStatus(model.HTTP_FORBIDDEN)
				return
			}
		}
	}
}

