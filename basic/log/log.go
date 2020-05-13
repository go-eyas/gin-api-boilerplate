package log

import (
	"github.com/go-eyas/toolkit/log"
)

func Init(conf *log.LogConfig) {
	printCaller = conf.Caller
	log.Init(conf)
	SugaredLogger = log.SugaredLogger
	Logger = log.Logger
}
