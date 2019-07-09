package util

// 错误码
const (
	CodeSuccess           = 0      // 操作成功
	CodeValidaFail        = 100100 // body 数据结构验证失败
	CodeParamFail         = 100101 // url 参数错误
	CodeJSONFail          = 100102 // JSON 解析错误
	CodeUserVerifyFail    = 100103 // 用户验证失败
	CodeDBQueryFail       = 100200 // 数据查询失败
	CodeDBConnectFail     = 100201 // 数据库连接失败
	CodePermissionDefined = 100300 // 权限不足
	CodeUnknowError       = 999999 // 未知错误
	CodeJSONPaseFail      = 100900 // json 解析失败
)
