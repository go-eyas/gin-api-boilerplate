package config

import "github.com/BurntSushi/toml"

var defaultConfig = `
debug = true # 是否启用调试
runtime = "runtime" # 运行文件目录

[server]
addr = ":9000" # 监听地址
baseUrl = "http://api.example.com" # 最终部署后访问该服务的地址

# 日志配置
[log]
level = "debug" # 日志级别
path = "logs" # 保存日志路径
console = true # 控制台是否输出日志
maxAge = 15 # 保存多少天的日志
fileHour = 1 # 多少小时分割一次日志
caller = false # 是否打印行号

[cors]
origin = ["*"]
methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"]
headers = ["Content-Length", "Content-Type", "Origin"]
credentials = true
maxAge = 24 # 单位: 小时
`

func init() {
	_, err := toml.Decode(defaultConfig, Conf)
	if err != nil {
		panic("default config synctax error")
	}
}
