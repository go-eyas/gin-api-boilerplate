package handler

import (
	"api/util"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	util.R(ctx).OK("")
}
