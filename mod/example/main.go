package example

import (
  "api/main/config"
  "basic/api"
)

func Init(conf *config.Config) {
  api.Register(Route)
}
