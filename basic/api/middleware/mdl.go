package middleware

import (
	"basic/config"
	"time"

	"github.com/go-eyas/toolkit/log"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// Common 全局通用的中间件
func Common(r *gin.Engine, conf *config.Config) {
	// 请求日志
	r.Use(Ginzap(log.Logger, true))
	// 错误恢复
	r.Use(RecoveryWithZap(log.Logger, false))
	// 错误处理
	r.Use(ErrorMiddleware(log.SugaredLogger))

	// Cors
	corsConf := conf.Cors
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     corsConf.Origin,
		AllowMethods:     corsConf.Methods,
		AllowHeaders:     corsConf.Headers,
		AllowCredentials: corsConf.Credentials,
		AllowOriginFunc: func(origin string) bool {
			for _, host := range corsConf.Origin {
				if origin == host || host == "*" {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	}))
}
