/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package model

import (
	"time"
)

// 配置文件结构体
type Config struct {
	AppId         string
	Port          int
	Mode          string
	RunMode       string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration

	LogRoot       string
	LogPath       string
	Color         bool
	AuthMethod    string
	AuthKey       string

	// jjauth key
	JJAuthKey     string

	Cluster		  []string

	DBJJGo       string
	DBMysite     string
	DBBlog       string

	MiddleWare   map[string]string

	Domain       string
}


// 白名单
type WhiteList struct {
	Refer    []string `json:"refer"`
	Host     []string `json:"host"`
	IP       []string `json:"ip"`
}

// 黑名单
type BlackList struct {
	Refer    []string `json:"refer"`
	Host     []string `json:"host"`
	IP       []string `json:"ip"`
}

// 更新日志
type ChangeLog struct {
	Version     string `json:"version"`
	Build       string `json:"build"`
	Description string `json:"description"`
}

// 版本信息
type JJGoVersion struct {
	Version 	string `json:"version"`
	BuildDate   string `json:"build_date"`
	JJCLI       string `json:"jjcli"`
	JJLog       string `json:"jjlog"`
	APIServer   string `json:"api_server"`
}