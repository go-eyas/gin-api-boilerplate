package config

import (
	"time"
	"toolkit/config"
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
	Log struct {
		Path         string `default:"logs/"`
		Console      bool   `default:"true"`
		MaxDay       int    `default:"15"`
		FileHour     int    `default:"1"`
		Caller       bool   `default:"false"`
		MaxAge       time.Duration
		RotationTime time.Duration
	}

	// redis 配置
	Redis struct {
		Cluster   bool     `default:"false"`
		Addrs     []string `default:"['127.0.0.1:6379']"`
		Password  string
		DB        int    `default:"0"`
		Namespace string `default:"minikv4"`
	}
}

// Conf 配置项实例
var Conf = &Config{}

func init() {
	config.Init("config", Conf)

	Conf.Log.Path = Conf.Runtime + "/" + Conf.Log.Path
	Conf.Log.MaxAge = time.Hour * 24 * time.Duration(Conf.Log.MaxDay)
	Conf.Log.RotationTime = time.Hour * time.Duration(Conf.Log.FileHour)
}
