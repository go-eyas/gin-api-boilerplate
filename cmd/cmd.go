package cmd

import (
	"api/config"
	"api/log"
	"api/srv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var AppName = "API"
var AppVersion = "1.0.0"
var Tag = ""
var conf *config.Config

var rootCmd = &cobra.Command{
	Use:   AppName,
	Short: AppName + " is a Golang Gin api example",
	Long: AppName + ` is a Golang Gin out of box api example:
* logs: base on zap
* command line interface tool
* database: base on gorm
* database migration
* config: base on configor
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		apiCMD(conf)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + AppName,
	Long:  `All software has versions. This is ` + AppName + `'s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(AppName + " v" + AppVersion + Tag)
	},
}

func init() {
	// 初始化配置项
	conf = config.Conf

	// 创建运行目录
	os.MkdirAll(conf.Runtime, os.ModePerm)

	// 初始化日志
	log.Init(conf)

	// 初始化客户端
	srv.Init(conf)

	// 添加命令行工具
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(apiCmd)
}

// Execute 启动命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 关闭客户端
	defer srv.Close()
}
