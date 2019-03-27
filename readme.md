# Gin 开箱即用项目模板

用户 API 开发的开箱即用模板

## 使用的开源库

* github.com/spf13/cobra 命令行工具
* github.com/jinzhu/gorm 数据库 orm 工具
* gopkg.in/gormigrate.v1 为 gorm 定制的数据库迁移工具
* github.com/jinzhu/configor 配置工具，支持 TOML、json、yaml、shell 环境变量读取配置值
* github.com/gin-gonic/gin Gin http web 框架
* github.com/uber-go/zap 性能极高的日志处理库

## 已集成功能特性

* 数据库
* 数据库迁移
* 日志
* 命令行工具
* 配置项，支持 OML、json、yaml、shell 环境变量 读取
* web 框架
* 工具函数
  + 错误码
  + http 回应封装
    ```
    Resp(ctx).OK(gin.H{"hello": "world"})

    // 将会输出
    {
      "data": {
          "hello": "world"
      },
      "msg": "ok",
      "status": 0
    }
    ```
  + 错误码

#### web 中间件

* gin.Recovery 错误处理，即使应用异常后不会导致程序退出
* ErrorMiddleware 错误处理，产生错误后，把错误信息回应到http请求，支持设置错误信息，错误码，数据
* cors 跨域，基于 github.com/gin-contrib/cors ，已配置好，只需要改配置文件就好

## 使用

1. git clone --depth=1 https://github.com/go-eyas/gin-api-boilerplate
2. 运行脚本 `./build.sh dev`, 将会直接下载依赖并运行
    + 如果您的IDE是vscode，请按F5直接运行
    + 如果需要命令行运行，请先将目录切换到 `build/`，然后 `go run ../main.go`，不然配置文件用不了

#### `./build.sh` 脚本使用

```sh
$ ./build.sh

USAGE:
  ./run.sh <command>

COMMANDS:
  docker [version]            构建项目成 docker 镜像
  build                       构建适用于当前操作系统的可执行文件
  dev                         启动开发环境，自动重启，自动生成文档
  help                        打印帮助信息
```


## 命令行工具使用

```sh
# version
$ go run main.go version
API v1.0.0

# help, 
$ go run main.go help
API is a Golang Gin out of box api example:
* logs: base on zap
* command line interface tool
* database: base on gorm
* database migration
* config: base on configor

Usage:
  API [flags]
  API [command]

Available Commands:
  api         Start HTTP API Server
  help        Help about any command
  migrate     migrate database
  version     Print the version number of API

Flags:
  -h, --help   help for API

Use "API [command] --help" for more information about a command.

# api, start http server
$ go run main.go api
[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> api/handler.SayHello (6 handlers)
2019-03-27T18:10:32.955+0800    info    route/api.go:48 API Server Listening:
[GIN-debug] Listening and serving HTTP on
```