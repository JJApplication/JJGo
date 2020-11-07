/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package util

import (
	"crypto/sha256"
	"fmt"
)
// token的加密函数使用sha256
func TokenEncrypt(token string) string {
	sha := sha256.New()
	sha.Write([]byte(token))
	return fmt.Sprintf("%x", sha.Sum(nil))
}
