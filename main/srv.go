package main

import (
	"api/main/config"
	"basic/srv"
)

// 按需初始化，只初始化自己用得到的功能
func srvInit(conf *config.Config) {
	// init redis
	srv.RedisSrv.Init(&conf.Redis)

	// init db, gorm and xorm chose one
	srv.Gorm.Init(&conf.DB)
	// srv.Xorm.Init(&conf.DB)

	srv.EmailInit(&conf.Email)
}
