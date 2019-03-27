package route

import (
	"api/config"
	"api/handler"
	"api/log"
	"time"

	"api/route/middleware"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func APIRun(conf *config.Config) error {
	serveConf := conf.Server
	corsConf := conf.Cors
	if !serveConf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := gin.Default()

	// 通用中间件
	middleware.Common(api)

	// Cors
	api.Use(cors.New(cors.Config{
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

	routes(api)

	serverAddr := serveConf.Addr
	log.Logger.Infof("API Server Listening: %s", serverAddr)
	return api.Run(serverAddr)
}

func routes(router *gin.Engine) {
	router.GET("/", handler.SayHello)
}
