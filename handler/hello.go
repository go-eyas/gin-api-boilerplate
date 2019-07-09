package handler

import (
	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	panic(gin.H{
		"status": 123456,
		"msg":    "asdfdsfs",
		"data": gin.H{
			"demo": "hello world",
		},
	})

}
