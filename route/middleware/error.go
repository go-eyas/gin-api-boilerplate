package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var codeUnknowError = 999999

// ErrorMiddleware 捕获到在http处理时的错误
// 在 handler 和其他地方如果产生了 error 可直接panic，到这里统一处理，简化 if err != nil 之类的代码
// panic("text") => {msg: "text", code: 0, data: {}}
// panic(gin.H{"code": 0, "msg": "some error"}) => {与传入的数据一致，} code 默认999999，status 默认 400，msg 默认 unknow error
// panic(errors.New("some error")) => {msg: "some error", code: 999999, data: {}}
// panic(Struct{...}) => {msg: "unknow", code: 999999, data: {...struct 数据}}
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case gin.H:
					e := err.(gin.H)
					code := e["status"]
					if code == nil {
						code = codeUnknowError
					}
					msg := e["msg"]
					if msg == nil {
						msg = "unknow error"
					}
					statusV := e["status"]
					statusCode := http.StatusBadRequest
					if statusV != nil {
						statusCode = statusV.(int)
					}
					ctx.Abort() // 需要显式终止中间件，不然依然会往里层调用
					ctx.JSON(statusCode, gin.H{
						"status": code,
						"msg":    msg,
						"data":   &gin.H{},
					})
				case error:
					e := err.(error)
					ctx.Abort()
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status": codeUnknowError,
						"msg":    e.Error(),
						"data":   &gin.H{},
					})
				case string:
					e := err.(string)
					ctx.Abort()
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status": codeUnknowError,
						"msg":    e,
						"data":   &gin.H{},
					})
				default:
					ctx.Abort()
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status": codeUnknowError,
						"msg":    "unknow error",
						"data":   err,
					})
				}
			}
		}()
		ctx.Next()
	}
}
