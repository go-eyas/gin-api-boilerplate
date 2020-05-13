package model

import (
	"api/main/config"
	"basic/srv"
)

func Init(conf *config.Config) {
	// 自动建表
	if conf.Debug {
		srv.Gorm.Migrate(Product{}) // gorm
		// db.Xorm.Migrate(Product{}) // xorm
	}
}
