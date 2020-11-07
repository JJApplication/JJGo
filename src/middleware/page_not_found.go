/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/model"
	"jjgo/src/util"
)

// 默认的404处理中间件
func PageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errCode := c.Writer.Status()
		if errCode == 404 {
			util.JJResponse(
				c,
				model.HTTP_NOT_FOUND,
				// 因为想要输出错误页面信息
				"can't find what you are looking for!",
				"404 Not Found",
			)
		}
	}
}
