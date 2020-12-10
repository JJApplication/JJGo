/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

配置读取函数， 配置由INI文件生成
*/
package config

import (
	"errors"
	"jjgo/src/logger"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-ini/ini"
	"jjgo/src/model"
)

var JJGoConf model.Config

func init() {
	JJGoConf, _ = ReadConf(true)
}

func ReadConf(init bool) (model.Config, error) {
	var conf model.Config
	cwdPath , pathErr := os.Getwd()
	if pathErr != nil {
		return model.Config{}, errors.New("get config failed")
	}
	iniPath := path.Join(cwdPath, "conf" ,"jjgo.ini")

	cfg, err := ini.Load(iniPath)
	if err != nil {
		return model.Config{}, errors.New("read config failed")
	}
	// 配置输出到结构体中
	conf = model.Config{
		AppId: cfg.Section("server").Key("app_id").String(),
		Port: cfg.Section("server").Key("port").MustInt(8020),
		Mode: cfg.Section("server").Key("mode").String(),
		RunMode: cfg.Section("server").Key("run_mode").String(),
		ReadTimeout: time.Duration(cfg.Section("server").Key("read_timeout").MustInt(60)) * time.Second,
		WriteTimeout: time.Duration(cfg.Section("server").Key("write_timeout").MustInt(60)) * time.Second,
		IdleTimeout: time.Duration(cfg.Section("server").Key("idle_timeout").MustInt(60)) * time.Second,

		LogRoot: cfg.Section("log").Key("log_root").String(),
		LogPath: cfg.Section("log").Key("log_path").String(),

		AuthMethod: cfg.Section("auth").Key("auth_method").String(),
		AuthKey: cfg.Section("auth").Key("auth_key").String(),

		Cluster: strings.Fields(cfg.Section("cluster").Key("ports").String()),

		DBJJGo: cfg.Section("database").Key("db_jjgo").String(),
		DBMysite: cfg.Section("database").Key("db_mysite").String(),
		DBBlog: cfg.Section("database").Key("db_blog").String(),
	}
	if !init {
		confLogger := logger.InitLogger(conf.LogPath, conf.Mode)
		confLogger.Info("配置文件读取完毕")
	}

	return conf, nil
}
