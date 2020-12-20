/*
JJGo api design
landers
*/

package url

// urls for jjgo
const PREFIX_JJGO = "/rest/jjgo"

// 根据url组的规则 不应该带前缀
const  (
	JJGO_INDEX = ""
	JJGO_PUB = "/pub"
	JJGO_STATUS = "/status"
	JJGO_CHANGELOG  = "/changelog"
	JJGO_VERSION = "/version"
	JJGO_DEMO = "/demo"
	JJGO_SWAGGER = "swagger.json"
)

func URL2JJGO(raw string) string {
	return PREFIX_JJGO + raw
}
