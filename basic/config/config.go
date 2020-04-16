package config

import (
	"basic/log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-eyas/toolkit/config"
)

// Config 配置
type Config struct {
	// 启用调试
	Debug bool `env:"API_DEBUG"`
	// 运行目录
	Runtime string `default:"runtime"`

	// http 服务器配置
	Server struct {
		Addr    string `default:"0.0.0.0:9000" env:"API_ADDR"`
		BaseURL string
	}

	// cors 跨域配置
	Cors struct {
		Origin      []string `default:"[*]"`
		Methods     []string `default:"['GET','POST','PUT','DELETE','PATCH','OPTIONS']"`
		Headers     []string `default:"['Content-Length','Content-Type','Origin']"`
		Credentials bool     `default:"true"`
		MaxAge      int      `default:"24"`
	}

	// 数据库配置
	DB struct {
		Driver string `default:"mysql" env:"MK_DB_DRIVER"`
		URI    string `env:"MK_DB_URI"`
	}

	// 日志配置
	Log log.Config

	// redis 配置
	Redis struct {
		Cluster   bool     `default:"false"`
		Addrs     []string `default:"['127.0.0.1:6379']"`
		Password  string
		DB        int    `default:"0"`
		Namespace string `default:"minikv4"`
	}
}

var defaultConfigContent = `
debug = true # 是否启用调试
runtime = "runtime" # 运行文件目录

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

// Conf 配置项实例
var Conf = &Config{}

func init() {
	toml.Decode(defaultConfigContent, Conf)
	config.Init("api", Conf)

	Conf.Log.Path = Conf.Runtime + "/" + Conf.Log.Path
	Conf.Log.MaxAge = time.Hour * 24 * time.Duration(Conf.Log.MaxAge)
	Conf.Log.RotationTime = time.Hour * time.Duration(Conf.Log.RotationTime)
}
