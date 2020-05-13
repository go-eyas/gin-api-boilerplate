package main

import (
	"api/main/config"
	"basic/api"
	"basic/log"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/gin-contrib/cors"
)

// 运行 http 服务器
func httpRun(conf *config.Config) {
	// 开发环境，自动运行建表
	if conf.Debug {
		migrateRun(conf)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 启动服务
	corsConf := conf.Cors
	router := api.NewApi(
		cors.New(cors.Config{
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
		}),
	)

	api.SetRoutes(router)

	log.Infof("API Server Listening: %s", conf.Server.Addr)
	if err := router.Run(conf.Server.Addr); err != nil {
		log.Fatalf("run api error: %v", err)
		panic(err)
	}
}
