package util

import (
	"encoding/json"
)

func Assert(err error, msg interface{}) {
	if err != nil {
		panic(msg)
	}
}

func AssertParam(err error) {
	if err != nil {
		panic(&RData{
			Status: CodeParamFail,
			Msg:    "param error: " + err.Error(),
		})
	}
}

// StructToMap 把结构体转成map，key使用json定义的key
func StructToMap(v interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	bt, _ := json.Marshal(v)
	json.Unmarshal(bt, &data)

	return data
}
