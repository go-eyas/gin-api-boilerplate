package db

import (
	"api/config"
	"toolkit/db"
	"toolkit/log"

	"github.com/go-xorm/xorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3 需要cgo编译环境，如果真的需要sqlite3再取消这行注释
)

type XormClient struct {
	DB *xorm.Engine
}

// Xorm 用户初始化和清理数据库
var Xorm = &XormClient{}

// XDB 数据库实例
var XDB *xorm.Engine

// // 多个数据库实例
// var XDB2 *xorm.Engine

// Init 初始化数据库
func (d *XormClient) Init(conf *config.Config) {
	dbConf := conf.DB
	var err error

	// 如果有多个数据库，复制这段继续
	if dbConf.Driver != "" && dbConf.URI != "" {
		XDB, err = db.Xorm(&db.Config{
			Driver: dbConf.Driver,
			URI:    dbConf.URI,
			Debug:  conf.Debug,
			Logger: log.SugaredLogger,
		})
		if err != nil {
			log.Fatalf("initial database error: %v", err)
			panic(err)
		}
	}
}

// Close 关闭数据库连接
func (d *XormClient) Close() {
	if d.DB != nil {
		log.Info("正在关闭数据库连接...")
		d.DB.Close()
	}
}
