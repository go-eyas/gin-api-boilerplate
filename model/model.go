package model

import (
	"basic/config"
	"basic/srv/db"
)

func Init(conf *config.Config) {
	// 自动建表
	if conf.Debug {
		db.Gorm.Migrate(Product{}) // gorm
		// db.Xorm.Migrate(Product{}) // xorm
	}
}
