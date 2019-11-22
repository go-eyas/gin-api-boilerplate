package db

import (
	"basic/config"
	"github.com/go-eyas/toolkit/db"
	"basic/log"

	"github.com/jinzhu/gorm"
	// _ "toolkit/db/sqlite" // sqlite3 需要cgo编译环境，如果真的需要sqlite3再取消这行注释
)

type GormClient struct {
	DB *gorm.DB
}

// Gorm 用户初始化和清理数据库
var Gorm = &GormClient{}

// GDB 数据库实例
var GDB *gorm.DB

// // 多个数据库实例
// var GDB2 *gorm.DB

// Init 初始化数据库
func (d *GormClient) Init(conf *config.Config) {
	dbConf := conf.DB
	var err error

	// 如果有多个数据库，复制这段继续
	if dbConf.Driver != "" && dbConf.URI != "" {
		GDB, err = db.Gorm(&db.Config{
			Driver: dbConf.Driver,
			URI:    dbConf.URI,
			Debug:  conf.Debug,
			Logger: log.Logger,
		})
		if err != nil {
			log.Fatalf("initial database error: %v", err)
			panic(err)
		}
	}
}

// Close 关闭数据库连接
func (d *GormClient) Close() {
	if d.DB != nil {
		log.Info("正在关闭数据库连接...")
		d.DB.Close()
	}
}
