package db

import (
	"api/config"
	"api/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init(conf *config.Config) *gorm.DB {
	dbConf := conf.DB

	DB, err := gorm.Open(dbConf.Type, dbConf.URI)

	if err != nil {
		log.Logger.Fatalf("initial database error: %v", err)
		panic(err)
	}

	if dbConf.AutoMigrate {
		// 自动运行 migrate
		Migrate()
	}

	return DB
}
