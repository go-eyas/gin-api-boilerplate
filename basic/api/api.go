package api

import (
	"basic/api/middleware"
	"basic/log"
	"regexp"

	"github.com/gin-gonic/gin"
)

func New(mdls ...gin.HandlerFunc) *gin.Engine {
	api := gin.New()
	api.Use(mdls...)
	return api
}

func NewApi(mdls ...gin.HandlerFunc) *gin.Engine {
	defMdls := []gin.HandlerFunc{
		middleware.Ginzap(log.Logger, true, regexp.MustCompile("/*/*.(js|css|png|jpg|woff|tff|oet|html)?")),
		middleware.RecoveryWithZap(log.Logger, false),
		middleware.ErrorMiddleware(log.SugaredLogger),
	}
	return New(defMdls...)
}
