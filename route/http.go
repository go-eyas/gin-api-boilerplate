package route

import (
	"api/handler"
	"basic/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

// 路由定义
func Routes(router *gin.Engine) {
	// 默认路由到 public 目录
	box := packr.NewBox("../public")
	router.Use(middleware.AssetsNoRoute(&box))

	router.GET("/ping", handler.Ping)
}
