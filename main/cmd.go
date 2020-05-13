package main

import (
	"api/main/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:     "api",
	Aliases: []string{"server", "serve"},
	Short:   "启动 http 服务器",
	Run: func(cmd *cobra.Command, args []string) {
		if config.Conf.Server.Addr == "" {
			var addr string
			switch len(args) {
			case 2:
				addr = args[0] + ":" + args[1]
			case 1:
				addr = args[0]
				if !strings.Contains(addr, ":") {
					addr = "0.0.0.0:" + addr
				}
			}
			if addr != config.Conf.Server.Addr {
				config.Conf.Server.Addr = addr
			}
		}
		httpRun(config.Conf)
	},
}

// 自动初始化数据库命令
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"createdbtable"},
	Short:   "自动创建数据表",
	Run: func(cmd *cobra.Command, args []string) {
		migrateRun(config.Conf)
	},
}

func Cmd() *cobra.Command {
	var RootCmd = &cobra.Command{
		Use: "server",
		Run: apiCmd.Run,
	}

	RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "输出版本号",
		Long:  `All software has versions. This is ` + appName + `'s`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(
				appName + " v" + version +
					"\nGit Commit:" + gitCommit +
					"\nBuild Time:" + buildTime +
					"\nGo Version:" + goVersion,
			)
		},
	})

	RootCmd.AddCommand(apiCmd) // api run

	return RootCmd
}
