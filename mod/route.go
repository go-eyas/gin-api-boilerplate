package mod

import (
	"api/assets"
	"api/main/config"
	"basic/api/middleware"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// 默认路由到 public 目录
	router.Use(middleware.AssetsNoRoute("/", &assets.Public))

	// 文档
	if config.Conf.Debug {
		router.Use(middleware.Assets("/docs", &assets.Docs))
	}
}
