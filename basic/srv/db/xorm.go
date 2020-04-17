package db

import (
	"basic/config"
	"basic/log"

	"github.com/go-eyas/toolkit/db"

	"github.com/go-xorm/xorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3 需要cgo编译环境，如果真的需要sqlite3再取消这行注释
)

type XormClient struct {
	DB              *xorm.Engine
	existTableCache map[string]bool
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
			Logger: log.Logger,
		})
		if err != nil {
			log.Fatalf("initial database error: %v", err)
			panic(err)
		}
		d.DB = XDB
		d.existTableCache = make(map[string]bool)
	}
}

type MigrateModelX interface {
	TableName() string
}

// Migrate 自动建表，表存在时不做任何操作，所以不能自动更新表字段
func (d *XormClient) Migrate(models ...MigrateModelX) {
	if d.DB == nil {
		log.Error("database not init")
		return
	}
	for _, m := range models {
		if exist, err := d.DB.IsTableExist((m).TableName()); !exist && err == nil {
			d.DB.Sync2(m)
		}
	}
}

// CheckTable 检查是否存在表，如果不存在则自动新建
func (d *XormClient) CheckTable(name string, model interface{}) *xorm.Session {
	if !d.existTableCache[name] {
		if exist, err := d.DB.IsTableExist(name); !exist && err == nil {
			d.DB.Sync2(model)
		}
		d.existTableCache[name] = true
	}
	return d.DB.Table(name)
}

// SplitDB 自动分表，根据 TableName 返回值获取表名
func (d *XormClient) SplitDB(model MigrateModelX) *xorm.Session {
	tbName := model.TableName()
	return d.CheckTable(tbName, model)
}

// Close 关闭数据库连接
func (d *XormClient) Close() {
	if d.DB != nil {
		log.Info("正在关闭数据库连接...")
		d.DB.Close()
	}
}
