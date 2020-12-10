/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

日志记录器
*/
package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"jjgo/src/logger"
	"jjgo/src/model"
)

func JJGoLog() gin.HandlerFunc {
	formatter := defaultLogFormatter

	out := logger.JJGoLogger.LogWriter

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		param := model.LogFormatterParams{
			Request: c.Request,
			IsTerm:  isTerm,
			Keys:    c.Keys,
		}

		param.TimeStamp = time.Now()

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(1 << 0).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		_, _ = fmt.Fprint(out, formatter(param))
		}
	}


var defaultLogFormatter = func(param model.LogFormatterParams) string {
	return fmt.Sprintf("[JJGo] %v | %v %s | %v | [%s] |Refer: %s | HOST: %s | ERRS:%s\n",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.StatusCode,
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Referer(),
		param.Request.Host,
		param.ErrorMessage,
	)
}