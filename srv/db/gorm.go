package db

import (
	"api/config"
	"api/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3 需要cgo编译环境，如果真的需要sqlite3再取消这行注释
)

type dbLogger struct{}

func (dbLogger) Print(v ...interface{}) {
	log.Logger.Info(gorm.LogFormatter(v...)...)
}

var dbl = &dbLogger{}

type dataBaseClient struct {
	DB *gorm.DB
}

// DataBase 用户初始化和清理数据库
var DataBase = &dataBaseClient{}

// DB 数据库实例
var DB *gorm.DB

// // 多个数据库实例
// var DB2 *gorm.DB

// Init 初始化数据库
func (d *dataBaseClient) Init(conf *config.Config) {
	dbConf := conf.DB
	var err error

	// 如果有多个数据库，复制这段继续
	if dbConf.Driver != "" && dbConf.URI != "" {
		DB, err = gorm.Open(dbConf.Driver, dbConf.URI)
		if err != nil {
			log.Fatalf("initial database error: %v", err)
			panic(err)
		}
		DB.LogMode(true)
		DB.SetLogger(dbl)
	}
}

// Close 关闭数据库连接
func (d *dataBaseClient) Close() {
	if d.DB != nil {
		log.Info("正在关闭数据库连接...")
		d.DB.Close()
	}
}
