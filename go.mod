module api

go 1.12

require (
	github.com/99designs/gqlgen v0.8.2
	github.com/gin-contrib/cors v0.0.0-20190301062745-f9e10995c85a
	github.com/gin-gonic/gin v1.4.0
	github.com/go-eyas/toolkit v1.0.1
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-xorm/xorm v0.7.4
	github.com/gobuffalo/packr v1.25.0
	github.com/jinzhu/configor v1.1.0
	github.com/jinzhu/gorm v1.9.10
	github.com/rs/xid v1.2.1
	github.com/spf13/cobra v0.0.3
	github.com/uber-go/zap v1.9.1 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/vektah/gqlparser v1.1.2
	go.uber.org/zap v1.10.0
	gopkg.in/gormigrate.v1 v1.4.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	toolkit v0.0.0-00010101000000-000000000000
)

replace toolkit => github.com/go-eyas/toolkit v1.0.1
