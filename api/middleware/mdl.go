package middleware

import (
	"api/config"
	"api/log"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gobuffalo/packr"
)

// Common 全局通用的中间件
func Common(r *gin.Engine, conf *config.Config) {
	// static
	box := packr.NewBox("../../public")
	r.Use(Assets("/", &box))

	
	r.Use(Ginzap(log.RequestLogger, time.RFC3339, false))
	r.Use(RecoveryWithZap(log.RequestLogger, false))
	// r.Use(gin.Recovery())
	r.Use(ErrorMiddleware(log.Logger))

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
