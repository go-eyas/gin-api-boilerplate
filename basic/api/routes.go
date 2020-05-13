package api

import (
  "github.com/gin-gonic/gin"
)

type RegisterHandler func(engine *gin.Engine)

var RegisterHandles = []RegisterHandler{}

func Register(handle RegisterHandler) {
  RegisterHandles = append(RegisterHandles, handle)
}

func SetRoutes(router *gin.Engine) {
  for _, h := range RegisterHandles {
    h(router)
  }
}