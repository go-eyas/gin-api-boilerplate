package main

import (
	"api/model"
	"api/route"
	"api/service"
	"basic/api"
	"basic/cmd"
	"basic/config"
	"basic/log"
	"basic/srv"
	"fmt"
	"os"
	"runtime"
)

var appName = "server"
var description = appName + ` is a Golang Gin out of box api example:
* logs: base on zap
* command line interface tool
* database: base on gorm
* database migration
* config: base on configor
	`
var version = "1.0.0"
var gitCommit = "unknow"
var buildTime = "unknow"
var goVersion = runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH

func main() {
	// 初始化配置项
	conf := config.Conf

	// 创建运行目录
	err := os.MkdirAll(conf.Runtime, os.ModePerm)
	if err != nil {
		fmt.Println("运行目录初始化失败")
		panic(err)
	}

	// 初始化日志
	conf.Log.Name = appName
	log.Init(&config.Conf.Log)

	// 初始化客户端
	srv.Init(conf)

	// 初始化服务
	service.Init(conf)

	// api 初始化
	api.Routes = route.Routes

	// 数据模型
	cmd.ModelInit = model.Init

	// 运行命令
	cmd.Execute(&cmd.App{
		Name:        appName,
		Description: description,
		Version:     version,
		GitCommit:   gitCommit,
		BuildTime:   buildTime,
		GoVersion:   goVersion,
	})
}
