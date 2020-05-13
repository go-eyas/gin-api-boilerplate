package example

import (
  "github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
  router.GET("/example", func(c *gin.Context) {})
}
