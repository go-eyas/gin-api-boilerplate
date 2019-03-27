package handler

import (
	"api/util"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	util.Resp(ctx).OK(gin.H{
		"hello": "world",
	})

}
