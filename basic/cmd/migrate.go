package cmd

import (
	"basic/config"
	"basic/log"
	"basic/srv/db"
	"encoding/json"

	"github.com/spf13/cobra"
)

var ModelInit = func(*config.Config) {}

// 自动初始化数据库命令
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"createdbtable"},
	Short:   "自动创建数据表",
	Run: func(cmd *cobra.Command, args []string) {
		migrateHandler(config.Conf)
	},
}

func migrateHandler(conf *config.Config) {
	conf.Debug = true
	log.Info("正在初始化数据表...")
	if db.GDB == nil && db.XDB == nil {
		dbConf, _ := json.Marshal(conf.DB)
		log.Errorf("数据表初始化失败，数据库未初始化，请检查您的数据库配置: %s", string(dbConf))
		return
	}
	ModelInit(conf)
	log.Info("数据表初始化完成")
}
