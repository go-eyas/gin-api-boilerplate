package handler

import (
	"basic/util"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	util.R(c).OK("PONG")
}

func SayHello(ctx *gin.Context) {
	panic(gin.H{
		"status": 123456,
		"msg":    "asdfdsfs",
		"data": gin.H{
			"demo": "hello world",
		},
	})

}
