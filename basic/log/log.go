package log

import (
  "github.com/go-eyas/toolkit/log"
)

type Config = log.LogConfig

func Init(conf *Config)  {
  printCaller = conf.Caller
  log.Init(conf)
  Logger = log.SugaredLogger
}