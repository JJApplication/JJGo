/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

Python库调用
采用传入参数的方式调用celery任务
mail server更名为jjmail
ID: app_jjmail
*/
package jjmail

import (
	"encoding/json"
	"fmt"
	"jjgo/src/url"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/logger"
	"jjgo/src/middleware"
	"jjgo/src/model"
	"jjgo/src/model/database"
	"jjgo/src/model/entity"
	"jjgo/src/util"
)

var ApiJJMail *gin.RouterGroup

func JJMail(r *gin.Engine) {
	info := func(c *gin.Context) {
		util.JJResponse(
			c,
			model.HTTP_STATUS_OK,
			"this's API for JJMail",
			map[string]string{"api": "[jjmail]", "path": c.FullPath()},
		)
	}
	ApiJJMail = r.Group(url.PREFIX_JJMAIL)
	{
		ApiJJMail.GET(url.JJMAIL_INDEX, info)
		ApiJJMail.POST(url.JJMAIL_STATUS, jjmailStatus)
		ApiJJMail.PUT(url.JJMAIL_SUB_BLOG, middleware.TokenAuth(), jjmailSubBlog)
		ApiJJMail.DELETE(url.JJMAIL_UNSUB_BLOG, jjmailUnsubBlog)
		ApiJJMail.GET(url.JJMAIL_UNSUB_BLOG, jjmailUnsubBlog)
		// 用于发送指定消息
		ApiJJMail.POST(url.JJMAIL_SEND, middleware.TokenAuth(), jjmailSend)
		ApiJJMail.PUT(url.JJMAIL_REPLY, middleware.TokenAuth(), jjmailReply)
		ApiJJMail.PUT(url.JJMAIL_SUB_MGEK, middleware.TokenAuth(), jjmailSubMgek)
		ApiJJMail.DELETE(url.JJMAIL_UNSUB_MGEK, jjmailUnsubMgek)
		ApiJJMail.GET(url.JJMAIL_UNSUB_MGEK, jjmailUnsubMgek)
	}
}

// 判断是否在运行
func jjmailStatus(c *gin.Context) {
	flag := checkStatus()
	if flag {
		util.JJResponse(c, model.HTTP_STATUS_OK, "jjmail is running",
			model.SUCCESS,
			)
	}else {
		util.JJResponse(c, model.HTTP_STATUS_OK, "jjmail is running",
			model.DEAD,
		)
	}
}

// 订阅博客
func jjmailSubBlog(c *gin.Context) {
	// 获取参数中的邮箱地址
	body := entity.JJMailAddress{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	// 获取数据库句柄
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	var posts []database.Blog
	err = db.FindAll(&posts).Error
	if err != nil  || len(posts) <= 5 {
		logger.JJGoLogger.Error("数据库读取失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
		return
	}
	defer db.Close()
	res, _ := json.Marshal(posts[len(posts)-5:])
	address := body.Address
	cwd, _ := os.Getwd()
	lib_python := path.Join(cwd, "script", "python", "lib_jjmail.py")
	opt, err := exec.Command("python3", lib_python, address, "blog", string(res[:])).Output()
	if err != nil {
		logger.JJGoLogger.Error("执行jjmail脚本失败, blog", err)
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
		return
	}

	if string(opt[:]) == "" || len(opt) == 0 {
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			model.OK)
	}else {
		logger.JJGoLogger.Error("执行jjmail脚本结果, blog", string(opt[:]))
		logger.JJGoLogger.Info(fmt.Sprintf("脚本执行异常, 订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
	}
}

// blog取消订阅
func jjmailUnsubBlog(c *gin.Context) {
	// 获取参数中的邮箱地址
	if c.Request.Method == "GET" {
		address := c.Query("address")
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			address)
		return
	}
	body := entity.JJMailAddress{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	address := body.Address
	// 从redis中删除邮箱
	logger.JJGoLogger.Info(fmt.Sprintf("blog邮件取消订阅: %s", address))
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		body)
	return
}

// 自动回复
func jjmailReply(c *gin.Context) {
	// 获取参数中的邮箱地址
	body := entity.JJMailAddress{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	address := body.Address
	cwd, _ := os.Getwd()
	lib_python := path.Join(cwd, "script", "python", "lib_jjmail.py")
	opt, err := exec.Command("python3", lib_python, address, "reply").Output()
	if err != nil {
		logger.JJGoLogger.Error("执行jjmail脚本失败, reply", err)
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
		return
	}

	if string(opt[:]) == "" || len(opt) == 0 {
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			model.OK)
	}else {
		logger.JJGoLogger.Info(fmt.Sprintf("脚本执行异常, 订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
	}
}

// mgek订阅
func jjmailSubMgek(c *gin.Context) {
	// 获取参数中的邮箱地址
	body := entity.JJMailAddress{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	address := body.Address
	cwd, _ := os.Getwd()
	lib_python := path.Join(cwd, "script", "python", "lib_jjmail.py")
	opt, err := exec.Command("python3", lib_python, address, "mgek").Output()
	if err != nil {
		logger.JJGoLogger.Error("执行jjmail脚本失败, mgek", err)
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
		return
	}

	if string(opt[:]) == "" || len(opt) == 0 {
		logger.JJGoLogger.Info(fmt.Sprintf("订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			model.OK)
	}else {
		logger.JJGoLogger.Info(fmt.Sprintf("脚本执行异常, 订阅邮箱: %s", address))
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
	}
}

// mgek取消订阅
func jjmailUnsubMgek(c *gin.Context) {
	// 获取参数中的邮箱地址
	if c.Request.Method == "GET" {
		address := c.Query("address")
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			address)
		return
	}
	body := entity.JJMailAddress{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	address := body.Address
	// 从redis中删除邮箱
	logger.JJGoLogger.Info(fmt.Sprintf("mgek邮件取消订阅: %s", address))
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		model.OK)
	return
}

// 通用邮件发送服务
// 考虑安全性 暂时不实现
func jjmailSend(c *gin.Context) {
	// 获取参数中的邮箱地址
	body := entity.JJMailAddressData{}
	err := c.BindJSON(&body)
	if err != nil {
		// body读取失败
		logger.JJGoLogger.Error("读取请求体失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_PARSE_BODY_FAILED,
			model.BAD)
		return
	}
	address := body.Address
	data := body.Data
	cwd, _ := os.Getwd()
	lib_python := path.Join(cwd, "script", "python", "lib_jjmail.py")
	opt, err := exec.Command("python3", lib_python, address, "send", data).Output()
	if err != nil {
		logger.JJGoLogger.Error("执行jjmail脚本失败, send", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
		return
	}

	if string(opt[:]) == "" || len(opt) == 0 {
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
			model.OK)
	}else {
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_INNER_ERR,
			model.BAD)
	}
}