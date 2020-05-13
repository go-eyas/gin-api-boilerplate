package main

import (
	"api/main/config"
	"api/mod"
	"api/service"
	"basic/log"
	"fmt"
	"os"
)

// 运行前的准备工作

// 业务逻辑开始前，所有命令都会执行
func beforRun() {
	// 解析配置文件
	conf, err := LoadConfig(appName)
	if err != nil {
		fmt.Println("配置文件解析失败")
		panic(err)
	}

	// 创建运行目录
	err = os.MkdirAll(conf.Runtime, os.ModePerm)
	if err != nil {
		fmt.Println("运行目录初始化失败")
		panic(err)
	}

	// 初始化日志
	conf.Log.Name = appName
	log.Init(&conf.Log)
}

// 业务逻辑初始化
func serviceInit() {
	conf := config.Conf

	// 初始化客户端
	srvInit(conf)

	// 初始化服务
	service.Init(conf)

	// 初始化模块
	mod.Init(conf)
}
