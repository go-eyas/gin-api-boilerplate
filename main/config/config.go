package config

import (
	"basic/srv"

	"github.com/go-eyas/toolkit/log"
)

type Config struct {
	// 启用调试
	Debug bool `env:"API_DEBUG"`
	// 运行目录
	Runtime string `default:"runtime"`

	// http 服务器配置
	Server struct {
		Addr    string `env:"API_ADDR"`
		BaseURL string
	}

	// cors 跨域配置
	Cors struct {
		Origin      []string
		Methods     []string
		Headers     []string
		Credentials bool
		MaxAge      int
	}

	// 日志配置
	Log log.LogConfig

	Redis srv.RedisConfig

	DB srv.DBConfig

	Email srv.EmailConfig
}

// Conf 全局配置
var Conf = &Config{}
