package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
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
		Level   string `default:"info"`
		Path    string `default:"logs/app.log"`
		Console bool   `default:"true"`
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

var configFileName = "config"

// 自动搜索配置文件 config.xxx 并自动加载，如果配置文件不存在，使用默认配置
// 支持三种配置文件格式
// 并支持 环境变量覆盖配置文件
func init() {
	exts := []string{"toml", "json", "yml"}
	env := os.Getenv("CONFIG_ENV")
	if env == "" {
		env = "local"
	}

	filelist := []string{}

	for _, ext := range exts {
		filelist = append(filelist,
			configFileName+"."+ext,
			configFileName+"."+env+"."+ext,
			"../"+configFileName+"."+ext,
			"../"+configFileName+"."+env+"."+ext,
		)
	}

	hasFile := false

	for _, f := range filelist {
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			hasFile = true
			load(f)
		}
	}

	if !hasFile {
		print("没有找到配置文件: config.toml, config.json, config.yml, 使用默认配置启动")
	}

	// runtime 目录问题
	Conf.Log.Path = Conf.Runtime + "/" + Conf.Log.Path
	if Conf.DB.Driver == "sqlite3" {
		Conf.DB.URI = Conf.Runtime + "/" + Conf.DB.URI
	}

	if Conf.Server.BaseURL == "" {
		Conf.Server.BaseURL = Conf.Server.Addr
	}
}

// Load 加载解析配置文件
func load(file string) *Config {
	err := configor.Load(Conf, file)

	if err != nil {
		fmt.Printf("skip config file: %s, beacuse error: %v", file, err)
	}
	return Conf
}
