package route

import (
	"api/assets"
	"api/handler"
	"basic/api/middleware"

	"github.com/gin-gonic/gin"
)

// 路由定义
func Routes(router *gin.Engine) {
	// 默认路由到 public 目录
	router.Use(middleware.AssetsNoRoute(&assets.Public))

	router.GET("/ping", handler.Ping)
}
