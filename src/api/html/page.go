/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package html

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/url"
)

func JJGoIndex(r *gin.Engine) {
	index := r.Group(url.PREFIX_HTML)

	index.StaticFile(url.HTML_INDEX, "./static/index.html")
	index.StaticFile(url.HTML_DOCS, "./static/docs.html")
	index.StaticFile(url.HTML_CHNAGELOGS, "./static/changelogs.html")
	index.StaticFile(url.HTML_FAVICON, "./static/favicon.png")
	index.StaticFile(url.HTML_APPLE_ICON, "./static/apple-icon.png")
	index.StaticFile(url.HTML_JJGO_ICON, "./static/jjgo.png")
	index.StaticFile(url.HTML_MOUSE_ICON, "./static/pig.png")
	index.StaticFile(url.HTML_ROBOTS, "./static/robots.txt")
}
