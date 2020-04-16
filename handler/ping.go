package handler

import (
	"basic/util"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	util.R(c).OK("PONG")
}

/**
 * @api {get} /ping 检查服务状态是否正常
 * @apiGroup 2.其他
 *
 * @apiParamExample {curl} Curl:
 * curl $SERVICE/ping
 *
 * @apiSuccessExample Success-Response:
 * {
 *     "status": 0,
 *     "msg": "ok",
 *     "data": "PONG"
 * }
 *
 */
func SayHello(ctx *gin.Context) {
	panic(gin.H{
		"status": 123456,
		"msg":    "asdfdsfs",
		"data": gin.H{
			"demo": "hello world",
		},
	})

}
