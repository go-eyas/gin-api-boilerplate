package log

import (
	"idoc/config"

	"go.uber.org/zap"
)

func Init(conf *config.Config) {
	logConf := conf.Log
	cfg := &zap.Config{
		Level: logConf.Level,
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
}
