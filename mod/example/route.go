package example

import (
  "api/mod/example/handler"
  "github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
  router.GET("/ping", handler.Ping)
}
