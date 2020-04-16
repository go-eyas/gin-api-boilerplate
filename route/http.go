package route

import (
	"api/assets"
	"api/handler"
	"basic/api/middleware"
	"basic/config"

	"github.com/gin-gonic/gin"
)

// 路由定义
func Routes(conf *config.Config, router *gin.Engine) {
	// 默认路由到 public 目录
	router.Use(middleware.AssetsNoRoute("/", &assets.Public))

	// 文档
	if conf.Debug {
		router.Use(middleware.Assets("/docs", &assets.Docs))
	}

	router.GET("/ping", handler.Ping)
}
