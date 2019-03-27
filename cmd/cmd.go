package cmd

import (
	"api/config"
	"api/db"
	"api/log"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var AppName = "API"
var AppVersion = "1.0.0"
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
		fmt.Println(AppName + " v" + AppVersion)
	},
}

func init() {
	// init config
	conf = config.Load("config.toml")
	// init log
	log.Init(conf)

	// init database
	d := db.Init(conf)
	defer d.Close() // TODO: will exit now

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(migrateCmd)
}

// Execute 启动命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
