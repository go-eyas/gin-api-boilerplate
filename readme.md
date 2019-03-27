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