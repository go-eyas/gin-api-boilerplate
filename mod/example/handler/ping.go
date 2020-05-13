package handler

import (
  "basic/util"
  "github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
  util.R(c).OK("PONG")
}
