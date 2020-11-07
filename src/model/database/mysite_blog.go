/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package database

type Blog struct {
	Title   string `gorm:"column:title"   json:"title"`
	Url     string `gorm:"column:url"     json:"url"`
	Content string `gorm:"column:content" json:"content"`
}

type Post struct {
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url"   json:"url"`
	Post 	string `gorm:"column:post"  json:"post"`
	Date 	string `gorm:"column:date"  json:"date"`
	Tags 	string `gorm:"column:tags"  json:"tags"`
}

type Problems struct {
	Theme string `gorm:"column:theme" json:"theme"`
	Text  string `gorm:"column:text" json:"text"`
}

type Thoughts struct {
	Day  	string `gorm:"column:day" json:"day"`
	Thought string `gorm:"column:thought" json:"thought"`
}

type Message struct {
	Mess  string `gorm:"column:mess" json:"mess"`
	Time  string `gorm:"column:time" json:"time"`
}

type ViewsCount struct {
	View  string `gorm:"column:view" json:"view"`
	Count int    `gorm:"column:count" json:"count"`
}

type Music struct {
	Name 	string `gorm:"column:name" json:"name"`
	MusicId string `gorm:"column:music_id" json:"music_id"`
}


func (Blog) TableName() string {
	return "blog"
}

func (Post) TableName() string {
	return "post"
}

func (Problems) TableName() string {
	return "problems"
}

func (Thoughts) TableName() string {
	return "thoughts"
}

func (Message) TableName() string {
	return "dbmessage"
}

func (ViewsCount) TableName() string {
	return "views_count"
}

func (Music) TableName() string {
	return "music"
}