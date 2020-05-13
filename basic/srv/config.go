package srv

import (
	"github.com/go-eyas/toolkit/amqp"
	"github.com/go-eyas/toolkit/email"
	"github.com/go-eyas/toolkit/websocket"
)

type DBConfig struct {
	Driver string `default:"mysql" env:"DB_DRIVER"`
	URI    string `env:"DB_URI"`
	Debug  bool
}

type RedisConfig struct {
	Cluster   bool     `default:"false"`
	Addrs     []string `default:"[]"`
	Password  string
	DB        int    `default:"0"`
	Namespace string `default:"api"`
}

type EmailConfig = email.Config

type AmqpConfig = amqp.Config

type WSConfig = websocket.Config
