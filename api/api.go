package api

import (
	"api/config"
	"api/log"

	"api/api/middleware"

	"github.com/gin-gonic/gin"
)

// APIRun 启动 http api 服务
func APIRun(conf *config.Config) error {
	serveConf := conf.Server
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := gin.Default()

	// 通用中间件
	middleware.Common(api, conf)

	routes(api)

	serverAddr := serveConf.Addr
	log.Logger.Infof("API Server Listening: %s", serverAddr)
	return api.Run(serverAddr)
}
