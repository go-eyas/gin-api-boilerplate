package config

import (
	"github.com/jinzhu/configor"
)

// Config 配置
type Config struct {
	Server struct {
		Addr  string `default:"0.0.0.0:9000" env:"API_ADDR"`
		Debug bool   `default:"false" env:"API_DEBUG"`
	}
	Cors struct {
		Origin      []string `default:"[*]"`
		Methods     []string `default:"['GET','POST','PUT','DELETE','PATCH','OPTIONS']"`
		Headers     []string `default:"['Content-Length','Content-Type','Origin']"`
		Credentials bool     `default:"true"`
		MaxAge      int      `default:"24"`
	}
	DB struct {
		Type        string `default:"sqlite3"`
		URI         string `default:"database.db"`
		AutoMigrate bool   `default:"false"`
		Debug       bool   `default:"false" env:"API_DEBUG"`
	}
	Log struct {
		Level    string `default:"info"`
		Path     string `default:"logs/app.log"`
		Format   string `default:"console"`
		Simplify bool   `default:"false"`
		Debug    bool   `default:"false" env:"API_DEBUG"`
	}
}

var Conf = &Config{}

func Load(file string) *Config {
	configor.Load(Conf, file)
	return Conf
}
