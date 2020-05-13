package util

// 常量定义

// 错误码
const (
  CodeSuccess          = 0      // 操作成功
  CodeParamFail        = 100100 // 接口参数错误
  CodeJSONFail         = 100102 // JSON 解析错误
  CodePermissionDefine = 100200 // 权限不足
  CodeDBCreateFail     = 100103 // 数据库创建记录失败
  CodeDBUpdateFail     = 100103 // 数据库更新记录失败
  CodeDBQueryFail      = 100103 // 数据库查询记录失败
  CodeDBDeleteFail     = 100103 // 数据库删除记录失败
  CodeUnknowError      = 999999 // 未知错误
)
