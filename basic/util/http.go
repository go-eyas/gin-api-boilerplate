package util

import (
	"basic/log"
	"encoding/json"
	"time"

	"github.com/go-eyas/toolkit/http"
)

// Http http client
var Http = http.
	Type("json").                                      // 默认 json 发送
	Timeout(10 * time.Second).                         // 超时时间 10 秒
	UseRequest(func(req *http.Request) *http.Request { // 请求打日志
		var data interface{}
		if len(req.SuperAgent.Data) > 0 {
			data = req.SuperAgent.Data
		} else if len(req.SuperAgent.SliceData) > 0 {
			data = req.SuperAgent.SliceData
		} else if len(req.SuperAgent.RawString) > 0 {
			data = req.SuperAgent.RawString
		}
		body, _ := json.Marshal(data)
		log.Debugf("http 发送 %s %s header=%v data=%s", req.SuperAgent.Method, req.SuperAgent.Url, req.SuperAgent.Header, string(body))
		return req
	}).
	UseResponse(func(req *http.Request, res *http.Response) *http.Response { // 回应打日志
		log.Debugf("http 接收 %s %s %s", req.SuperAgent.Method, req.SuperAgent.Url, res.String())
		return res
	})

// HTTP alias 别名
var HTTP = Http
