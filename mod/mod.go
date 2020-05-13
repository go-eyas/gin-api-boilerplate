package mod

import (
	"api/main/config"
	"api/mod/example"
	"basic/api"
)

func Init(conf *config.Config) {
	api.Register(Route)

	example.Init(conf)

}
