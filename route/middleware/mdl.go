package middleware

import (
	"github.com/gin-gonic/gin"
)

// Common 全局通用的中间件
func Common(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(ErrorMiddleware())

}
