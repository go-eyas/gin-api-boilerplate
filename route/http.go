package route

import (
  "api/handler"
  "github.com/gin-gonic/gin"
)

// 路由定义
func Routes(router *gin.Engine) {
  router.GET("/ping", handler.Ping)
}