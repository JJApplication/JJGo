/*
App: JJGo
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-20
*/
package html

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/url"
)

// swagger UI页面
// 纯静态资源
func JJGoSwagger(r *gin.Engine) {
	swagger := r.Group(url.SWAGGER)

	swagger.StaticFile(url.SWAGGER_INDEX, "./swagger/web/index.html")
	swagger.StaticFile(url.SWAGGER_BUNDLE, "./swagger/web/swagger-ui-bundle.js")
	swagger.StaticFile(url.SWAGGER_PRESET, "./swagger/web/swagger-ui-standalone-preset.js")
	swagger.StaticFile(url.SWAGGER_UI_JS, "./swagger/web/swagger-ui.js")
	swagger.StaticFile(url.SWAGGER_UI_CSS, "./swagger/web/swagger-ui.css")
	swagger.StaticFile(url.SWAGGER_UI_ICON1, "./swagger/web/favicon-16x16.png")
	swagger.StaticFile(url.SWAGGER_UI_ICON2, "./swagger/web/favicon-32x32.png")
}
