package log

import (
	"api/config"
	"encoding/json"
	"os"
	"time"

	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// L 简单只能打日志
var L *zap.Logger

// Logger 可以格式化日志
var Logger *zap.SugaredLogger

// RequestLogger 用作接口请求记录
var RequestLogger *zap.Logger

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level         string
	Format        string
	Path          string
	Debug         bool
	OutputConsole bool
}

// Init 初始化日志
func Init(conf *config.Config) {
	logConf := conf.Log
	logPath := logConf.Path
	if err := os.MkdirAll(logPath+"/", os.ModePerm); err != nil {
		fmt.Println("init log path error.")
		panic(err)
	}

	// 应用运行日志
	L = New(&LoggerConfig{
		Level:         logConf.Level,
		Format:        "console",
		Debug:         conf.Debug,
		Path:          logPath + "/api.log",
		OutputConsole: logConf.Console,
	})
	Logger = L.Sugar()

	// 请求日志
	RequestLogger = New(&LoggerConfig{
		Level:         logConf.Level,
		Format:        "console",
		Debug:         conf.Debug,
		Path:          logPath + "/request.log",
		OutputConsole: false,
	})
}

// New 新建日志类型
func New(logConf *LoggerConfig) *zap.Logger {
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
	))

	var err error

	// file output
	var fileOut = zap.NewProductionConfig()
	if err = json.Unmarshal(rawJSON, &fileOut); err != nil {
		panic(err)
	}
	fileOut.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConf.Path,
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     30, // days
		LocalTime:  true,
	})

	// console output
	var consoleOut = zap.Config{}
	if err = json.Unmarshal(rawJSON, &consoleOut); err != nil {
		panic(err)
	}
	consoleWriter := zapcore.Lock(os.Stdout)

	// encoder
	var enc zapcore.Encoder
	enc = zapcore.NewConsoleEncoder(fileOut.EncoderConfig)

	// zap core
	cores := []zapcore.Core{
		zapcore.NewCore(enc, fileWriter, fileOut.Level),
	}

	if logConf.OutputConsole {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(consoleOut.EncoderConfig), consoleWriter, consoleOut.Level))
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core)

	defer logger.Sync()
	return logger
}

// Debugf 格式化日志
func Debugf(s string, v ...interface{}) {
	Logger.Debugf(s, v...)
}

// Infof 格式化日志
func Infof(s string, v ...interface{}) {
	Logger.Infof(s, v...)
}

// Warnf 格式化日志
func Warnf(s string, v ...interface{}) {
	Logger.Warnf(s, v...)
}

// Errorf 格式化日志
func Errorf(s string, v ...interface{}) {
	Logger.Errorf(s, v...)
}

// Fatalf 格式化日志
func Fatalf(s string, v ...interface{}) {
	Logger.Panicf(s, v...)
}

// Panicf 格式化日志
func Panicf(s string, v ...interface{}) {
	Logger.Panicf(s, v...)
}

// Debug 打日志
func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

// Info 打日志
func Info(v ...interface{}) {
	Logger.Info(v...)
}

// Warn 打日志
func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

// Error 打日志
func Error(v ...interface{}) {
	Logger.Error(v...)
}

// Panic 打日志
func Panic(v ...interface{}) {
	Logger.Panic(v...)
}
