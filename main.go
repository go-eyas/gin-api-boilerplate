package main

import (
	"api/route"
	"api/service"
	"basic/api"
	"basic/cmd"
	"basic/config"
	"basic/log"
	"basic/srv"
	"basic/srv/db"
	"fmt"
	"os"
	"runtime"
)

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
	conf.Log.Name = "server"
	log.Init(&config.Conf.Log)

	// 初始化客户端
	srv.Init(conf)

	// 初始化服务
	if db.GDB != nil {
		service.Init(conf)
	}

	// api 初始化
	api.Routes = route.Routes

	cmd.Execute(&cmd.App{
		Name:  "server",
		Short: "server is a Golang Gin api example",
		Description: ` is a Golang Gin out of box api example:
* logs: base on zap
* command line interface tool
* database: base on gorm
* database migration
* config: base on configor
	`,
		Version:   version,
		GitCommit: gitCommit,
		BuildTime: buildTime,
		GoVersion: goVersion,
	})
}
