/*
App: JJGo
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-20
*/
package html

import (
	"github.com/gin-gonic/gin"
)

// swagger UI页面
// 纯静态资源
func JJGoSwagger(r *gin.Engine) {
	swagger := r.Group("/swagger")

	swagger.StaticFile("", "./swagger/web/index.html")
	swagger.StaticFile("/static/swagger-ui-bundle.js", "./swagger/web/swagger-ui-bundle.js")
	swagger.StaticFile("/static/swagger-ui-standalone-preset.js", "./swagger/web/swagger-ui-standalone-preset.js")
	swagger.StaticFile("/static/swagger-ui.js", "./swagger/web/swagger-ui.js")
	swagger.StaticFile("/static/swagger-ui.css", "./swagger/web/swagger-ui.css")
	swagger.StaticFile("/static/favicon-16x16.png", "./swagger/web/favicon-16x16.png")
	swagger.StaticFile("/static/favicon-32x32.png", "./swagger/web/favicon-32x32.png")
}
