/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package model

// 格式响应
const (
	SUCCESS = "success"
	FAILED  = "failed"
	RUNNING = "running"
	DEAD    = "dead"

	// 针对部分接口响应格式
	OK      = "ok"
	BAD     = "bad"
)

// 返回msg提示信息
const (
	// 成功提示信息
	MSG_SUCCESS = "success"
	// 失败提示信息
	MSG_FAILED = "failed"
	// 预定义的token串
	MSG_FORBIDDEN_NEED_TOKEN = "forbidden to access, Token required"
	// jjtoken
	MSG_FORBIDDEN_NEED_JJTOKEN = "forbidden to access, JJToken required"
	// 解析body失败
	MSG_PARSE_BODY_FAILED = "parse body failed"
	// 解析query参数失败
	MSG_QUERY_PARAMS_FAILED = "query params failed"
	// 服务器内部错误
	MSG_INNER_ERR = "inner error"
)
