package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Responses base struct
type Responses struct {
	Ctx *gin.Context
}

func (r *Responses) responseError(data interface{}, err error) {
	ctx := r.Ctx
	switch data.(type) {
	case gin.H:
		ctx.JSON(http.StatusBadRequest, data)
	default:
		if data == nil {
			data = gin.H{}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": CodeUnknowError,
			"msg":    err.Error(),
			"data":   data,
		})
	}

}

// Response http 返回值包装，统一格式为 {code: 123, msg: "ok", data: {具体数据}}
func (r *Responses) Response(resp interface{}, err error) {
	ctx := r.Ctx
	if err != nil {
		r.responseError(resp, err)
		return
	}
	switch resp.(type) {
	case error:
		r.responseError(resp, err)
		return
	case gin.H:
		e := resp.(gin.H)
		code := e["status"]
		if code == nil {
			code = CodeSuccess
		}
		msg := e["msg"]
		if msg == nil {
			msg = "ok"
		}
		statusCode := http.StatusOK
		statusV := e["status"]
		if statusV != nil {
			statusCode = statusV.(int)
		}

		data := e["data"]
		if data == nil {
			data = resp
		}

		ctx.JSON(statusCode, gin.H{
			"status": code,
			"msg":    msg,
			"data":   data,
		})
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"status": CodeSuccess,
			"msg":    "ok",
			"data":   resp,
		})
	}

}

// Error 回应错误类型
func (r *Responses) Error(err error) {
	r.Response(nil, err)
}

// Bad 返回 400 状态码
func (r *Responses) Bad(data interface{}, err error) {
	r.Response(data, err)
}

// OK 返回 200 状态码
func (r *Responses) OK(data interface{}) {
	r.Response(gin.H{
		"status": CodeSuccess,
		"msg":    "ok",
		"data":   data,
	}, nil)
}

// Ok OK 的别名
func (r *Responses) Ok(data interface{}) {
	r.OK(data)
}

// Resp 创建返回对象
// 例子:
// resp := libs.Resp(ctx);
// resp.Ok("something")
func Resp(ctx *gin.Context) *Responses {
	return &Responses{Ctx: ctx}
}
