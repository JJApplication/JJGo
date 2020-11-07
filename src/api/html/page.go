/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package html

import (
	"github.com/gin-gonic/gin"
)

func JJGoIndex(r *gin.Engine) {
	index := r.Group("/")

	index.StaticFile("", "./static/index.html")
	index.StaticFile("/favicon.png", "./static/favicon.png")
	index.StaticFile("/apple-icon.png", "./static/apple-icon.png")
	index.StaticFile("/jjgo.png", "./static/jjgo.png")
	index.StaticFile("/pig.png", "./static/pig.png")
	index.StaticFile("/robots.txt", "./static/robots.txt")
}
