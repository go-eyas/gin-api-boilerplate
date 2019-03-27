package graphql

import (
	"api/graphql/g"
	"api/graphql/resolver"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// PlayGround dev tool
func PlayGround(title, path string) gin.HandlerFunc {
	h := handler.Playground(title, path)
	return gin.WrapH(h)
}

// Resolver graphql resovlers
func Resolver() gin.HandlerFunc {
	h := handler.GraphQL(g.NewExecutableSchema(g.Config{Resolvers: &resolver.Resolver{}}))
	return gin.WrapH(h)
}
