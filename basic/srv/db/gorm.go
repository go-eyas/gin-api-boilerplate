package db

import (
	"basic/config"
	"basic/log"

	"github.com/go-eyas/toolkit/db"

	"github.com/jinzhu/gorm"
	// _ "toolkit/db/sqlite" // sqlite3 需要cgo编译环境，如果真的需要sqlite3再取消这行注释
)

type GormClient struct {
	DB              *gorm.DB
	existTableCache map[string]bool
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
		d.DB = GDB
		d.existTableCache = make(map[string]bool)
	}
}

type MigrateModel interface {
	TableName() string
}

// Migrate 自动建表，表存在时不做任何操作，所以不能自动更新表字段
func (d *GormClient) Migrate(models ...MigrateModel) {
	for _, m := range models {
		if !d.DB.HasTable((m).TableName()) {
			d.DB.AutoMigrate(m)
		}
	}
}

// CheckTable 检查是否存在表，如果不存在则自动新建
func (d *GormClient) CheckTable(name string, model interface{}) *gorm.DB {
	if !d.existTableCache[name] {
		if !d.DB.HasTable(name) {
			d.DB.CreateTable(model)
		}
		d.existTableCache[name] = true
	}
	return d.DB.Table(name)
}

// SplitDB 自动分表，根据 TableName 返回值获取表名
func (d *GormClient) SplitDB(model MigrateModel) *gorm.DB {
	tbName := model.TableName()
	return d.CheckTable(tbName, model)
}

// Close 关闭数据库连接
func (d *GormClient) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
