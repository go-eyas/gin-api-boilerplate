package log

import (
	"api/config"
	"encoding/json"
	"os"

	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L *zap.Logger
var Logger *zap.SugaredLogger

func Init(conf *config.Config) {
	logConf := conf.Log
	rawJSON := []byte(fmt.Sprintf(`{
	  "level": "%s",
	  "encoding": "%s",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`,
		logConf.Level,
		logConf.Format,
		// logConf.Path,
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

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConf.Path,
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
		LocalTime:  true,
	})

	consoleWriter := zapcore.Lock(os.Stdout)

	var enc zapcore.Encoder
	if logConf.Format == "json" {
		enc = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	} else {
		enc = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}

	// core := zapcore.NewCore(enc, fileWriter, cfg.Level)
	core := zapcore.NewTee(
		zapcore.NewCore(enc, fileWriter, cfg.Level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), consoleWriter, cfg.Level),
	)

	logger := zap.New(core)

	// logger, err := cfg.Build()
	// if err != nil {
	// 	panic(err)
	// }
	defer logger.Sync()

	L = logger
	Logger = logger.Sugar()
}
