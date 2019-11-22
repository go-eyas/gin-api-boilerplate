package middleware

import (
	"basic/util"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)
type errLogger interface {
	Errorf(string, ...interface{})
}
// var codeUnknowError = 999999

// ErrorMiddleware 捕获到在http处理时的错误
// 在 handler 和其他地方如果产生了 error 可直接panic，到这里统一处理，简化 if err != nil 之类的代码
// panic("text") => {msg: "text", code: 0, data: {}}
// panic(gin.H{"code": 0, "msg": "some error"}) => {与传入的数据一致，} code 默认999999，status 默认 400，msg 默认 unknow error
// panic(errors.New("some error")) => {msg: "some error", code: 999999, data: {}}
// panic(Struct{...}) => {msg: "unknow", code: 999999, data: {...struct 数据}}
func ErrorMiddleware(logger errLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if data := recover(); data != nil {
				logger.Errorf("%v", data)
				debug.PrintStack()
				ctx.Abort()
				// res, err := util.RespError(data)
				r := util.R(ctx)
				r.Error(data)
				// if err != nil {
				// 	r.Error(util.RData{
				// 		Code: 500,
				// 		Msg: "unknow error",
				// 		Data: data,
				// 		Status: util.CodeUnknowError,
				// 	})
				// 	return
				// }
				// r.Error(res)
			}
		}()
		ctx.Next()
	}
}
