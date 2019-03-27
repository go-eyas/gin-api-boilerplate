package config

import (
	"github.com/jinzhu/configor"
)

// Config 配置
type Config struct {
	Server struct {
		Host string `default:"0.0.0.0" env:"IDOC_HOST"`
		Port string `default:"9000" env:"IDOC_PORT"`
	}
	DB struct {
		Type  string `default:"sqlite3"`
		URI   string `default:"database.db"`
		Debug bool   `default:"true"`
	}
	Log struct {
		Level string `default:"info"`
		Path  string `default:"logs"`
	}
}

var Conf = &Config{}

func Load(file string) *Config {
	configor.Load(Conf, file)
	return Conf
}
