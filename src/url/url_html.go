/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package url

// html url
const PREFIX_HTML = "/"

const(
	HTML_INDEX = ""
	HTML_DOCS = "docs"
	HTML_CHNAGELOGS = "changelogs"
	HTML_FAVICON = "/favicon.png"
	HTML_APPLE_ICON = "/apple-icon.png"
	HTML_JJGO_ICON = "/jjgo.png"
	HTML_MOUSE_ICON = "/pig.png"
	HTML_ROBOTS = "/robots.txt"
)

const(
	SWAGGER = "/swagger"
	SWAGGER_INDEX = ""
	SWAGGER_BUNDLE = "/static/swagger-ui-bundle.js"
	SWAGGER_PRESET = "/static/swagger-ui-standalone-preset.js"
	SWAGGER_UI_JS = "/static/swagger-ui.js"
	SWAGGER_UI_CSS = "/static/swagger-ui.css"
	SWAGGER_UI_ICON1 = "/static/favicon-16x16.png"
	SWAGGER_UI_ICON2 = "/static/favicon-32x32.png"
)

func URL2HTML(raw string) string {
	return PREFIX_HTML + raw
}