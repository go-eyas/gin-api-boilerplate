package main

import (
	. "api/main/config"
	"runtime"
	"time"

	"github.com/go-eyas/toolkit/config"
)

var appName = "api"
var description = appName + ` is a Golang Gin out of box api example:
* logs: base on zap
* command line interface tool
* database: base on gorm
* database migration
* config: base on configor
	`
var version = "1.0.0"
var gitCommit = "unknow"
var buildTime = "unknow"
var goVersion = runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH

func LoadConfig(name string) (*Config, error) {
	err := config.Init(name, Conf)

	// log
	Conf.Log.Path = Conf.Runtime + "/" + Conf.Log.Path
	Conf.Log.MaxAge = time.Hour * 24 * time.Duration(Conf.Log.MaxAge)
	Conf.Log.RotationTime = time.Hour * time.Duration(Conf.Log.RotationTime)

	// db
	Conf.DB.Debug = Conf.Debug
	return Conf, err
}
