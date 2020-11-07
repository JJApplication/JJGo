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
	ApiMysite = r.Group("/rest/mysite")
	{
		ApiMysite.GET("", info)
		ApiMysite.GET("/get_blog", mysiteBlog)
		ApiMysite.GET("/get_post", mysitePost)
		ApiMysite.GET("/get_problems", mysiteProblems)
		ApiMysite.GET("/get_thoughts", mysiteThoughts)
		ApiMysite.GET("/get_message", mysiteMessage)
		ApiMysite.GET("/get_music", mysiteMusic)
		ApiMysite.GET("/get_views", mysiteViews)
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "args required like ?url=",
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			database.Post{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "success", p)
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			[]database.Problems{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "success",
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			[]database.Thoughts{})
		return
	}

	util.JJResponse(c, model.HTTP_STATUS_OK, "success",
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			[]database.Message{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "success",
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			database.Message{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "success",
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
		util.JJResponse(c, model.HTTP_STATUS_OK, "failed",
			[]database.Music{})
		return
	}
	util.JJResponse(c, model.HTTP_STATUS_OK, "success",
		music)
	return
}