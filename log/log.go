package log

import (
	"encoding/json"
	"api/config"

	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger
var Logger *zap.SugaredLogger

func Init(conf *config.Config) {
	logConf := conf.Log
	rawJSON := []byte(fmt.Sprintf(`{
	  "level": "%s",
	  "encoding": "%s",
	  "outputPaths": ["stdout", "./%s"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`,
		logConf.Level,
		logConf.Format,
		logConf.Path,
	))
	var cfg zap.Config

	if !logConf.Debug {
		cfg = zap.NewProductionConfig()
	} else if logConf.Simplify {
		cfg = zap.Config{}
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	L = logger
	Logger = logger.Sugar()
}
