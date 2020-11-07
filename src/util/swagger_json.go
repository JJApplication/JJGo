/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"jjgo/src/model"
)

// 读取swagger文件
func SwaggerJson() interface{}{
	var swaggerJson interface{}
	cwd, _ := os.Getwd()
	jsonFile := path.Join(cwd, "swagger", "json", "swagger.json")
	data, err := ioutil.ReadFile(jsonFile)

	if err != nil {
		return model.Swagger{}
	}
	_ = json.Unmarshal(data, &swaggerJson)
	return swaggerJson
}