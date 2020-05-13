package mod

import (
	"api/main/config"
	"api/mod/example"
	"basic/api"
)

type Mod interface {
	Init(conf *config.Config)
	Route()
}

func Init(conf *config.Config) {
	api.Register(Route)

	example.Init(conf)

}

