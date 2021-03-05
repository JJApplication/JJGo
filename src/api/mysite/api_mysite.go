/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

仅支持部分api 如博客文章，问题，想法，音乐，访问量
*/
package mysite

import (
	"github.com/gin-gonic/gin"
	"jjgo/src/config"
	"jjgo/src/jjgorm"
	"jjgo/src/logger"
	"jjgo/src/model"
	"jjgo/src/model/database"
	"jjgo/src/url"
	"jjgo/src/util"
)

var ApiMysite *gin.RouterGroup

func Mysite(r *gin.Engine) {
	info := func(c *gin.Context) {
		util.JJResponse(
			c,
			model.HTTP_STATUS_OK,
			"this's API for Mysite",
			map[string]string{"api": "[mysite]", "path": c.FullPath()},
		)
	}
	ApiMysite = r.Group(url.PREFIX_MYSITE)
	{
		ApiMysite.GET(url.MYSITE_INDEX, info)
		ApiMysite.GET(url.MYSITE_GET_BLOG, mysiteBlog)
		ApiMysite.GET(url.MYSITE_GET_POST, mysitePost)
		ApiMysite.GET(url.MYSITE_GET_PROBLEMS, mysiteProblems)
		ApiMysite.GET(url.MYSITE_GET_THOUGHTS, mysiteThoughts)
		ApiMysite.GET(url.MYSITE_GET_MESSAGE, mysiteMessage)
		ApiMysite.GET(url.MYSITE_GET_MUSIC, mysiteMusic)
		ApiMysite.GET(url.MYSITE_GET_VIEWS, mysiteViews)
	}
}

// 博客列表
func mysiteBlog(c *gin.Context) {
	// 可选参数url 根据url取值
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	defer db.Close()

	var res []database.Blog
	err := db.FindAll(&res).Error

	if err != nil {
		logger.JJGoLogger.Error("数据库读取失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, "inner failed",
			[]database.Blog{})
		return
	}

	util.JJResponse(c, model.HTTP_STATUS_OK, "获取博客列表成功",
		res)
	return
}

// 博客文章
func mysitePost(c *gin.Context) {
	// 获取文章参数
	url := c.Query("url")
	if url == "" {
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_QUERY_PARAMS_FAILED,
			database.Post{})
		return
	}
	var p database.Post
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindFirstByStruct(database.Post{Url: url}, &p).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取文章数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			database.Post{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS, p)
	return
}

// 问题
func mysiteProblems(c *gin.Context) {
	var problems []database.Problems
 	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindAll(&problems).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取问题数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			[]database.Problems{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		problems)
	return
}

// 想法
func mysiteThoughts(c *gin.Context) {
	var thoughts []database.Thoughts
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindAll(&thoughts).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取想法数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			[]database.Thoughts{})
		return
	}

	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		thoughts)
	return
}

// 留言
func mysiteMessage(c *gin.Context) {
	var message []database.Message
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindAll(&message).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取留言数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			[]database.Message{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		message)
	return
}

// 访问量
func mysiteViews(c *gin.Context) {
	var views database.ViewsCount
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindFirstByStruct(database.ViewsCount{View: "all"}, &views).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取访问量数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			database.Message{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		views)
	return
}

// 音乐列表
func mysiteMusic(c *gin.Context) {
	var music []database.Music
	db := jjgorm.JJGormClient()
	db.Connect(config.JJGoConf.DBMysite)
	err := db.FindAll(&music).Error
	defer db.Close()
	if err != nil {
		logger.JJGoLogger.Error("获取歌单数据库失败", err)
		util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_FAILED,
			[]database.Music{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, model.MSG_SUCCESS,
		music)
	return
}