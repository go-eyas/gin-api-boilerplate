package api

import (
	// "api/graphql"
	"api/handler"

	"github.com/gin-gonic/gin"
)

// 路由定义
func routes(router *gin.Engine) {
	router.GET("/hello", handler.SayHello)

	// 如果你需要graphql的话
	// router.GET("/api/v1/graphql", graphql.PlayGround("Graphql API", "/api/v1/graphql"))
	// router.POST("/api/v1/graphql", graphql.Resolver())

}
