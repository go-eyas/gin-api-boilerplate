module api

go 1.12

require (
	basic v1.0.0
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.6.2
	github.com/go-eyas/toolkit v1.2.1
	github.com/gobuffalo/packr v1.30.1
	github.com/spf13/cobra v0.0.5
	xorm.io/xorm v1.0.1 // indirect
)

replace basic v1.0.0 => ./basic
