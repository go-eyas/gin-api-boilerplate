package db

import (
	// "github.com/jinzhu/gorm"
	"api/log"

	gormigrate "gopkg.in/gormigrate.v1"
)

// migrate options
var migrateOptions *gormigrate.Options = gormigrate.DefaultOptions

// define migrates
var migrateConfig []*gormigrate.Migration = []*gormigrate.Migration{
	// {
	// 	ID: "201903271804",
	// 	Migrate: func(db *gorm.DB) error {
	// 		type User struct {
	// 			gorm.Model
	// 			Username string `gorm:"column:username;unique_index;"`
	// 		}
	// 		return db.AutoMigrate(&User{}).Error
	// 	},
	// 	Rollback: func(db *gorm.DB) error {
	// 		return db.DropTable("users").Error
	// 	},
	// },
}

func Migrate() {
	if len(migrateConfig) > 0 {
		log.L.Info("database migrate start...")
		m := gormigrate.New(DB, migrateOptions, migrateConfig)
		if err := m.Migrate(); err != nil {
			log.Logger.Fatalf("could noe migrate database: %v", err)
			panic(err)
		}
		log.L.Info("migrate complate.")
	} else {
		log.L.Info("migrate config is empty, nothing to do.")
	}
}
