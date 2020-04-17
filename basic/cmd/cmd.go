package cmd

import (
	"basic/config"
	"basic/srv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type App struct {
	Name        string
	Short       string
	Description string
	Version     string
	GitCommit   string
	BuildTime   string
	GoVersion   string
}

var RootCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		apiCMD(config.Conf)
	},
}

func command(app *App) *cobra.Command {
	RootCmd.Use = app.Name
	RootCmd.Short = app.Short
	RootCmd.Long = app.Description

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "输出版本号",
		Long:  `All software has versions. This is ` + app.Name + `'s`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Name + " v" + app.Version + "\nGit Commit:" + app.GitCommit + "\nBuild Time:" + app.BuildTime + "\nGo Version:" + app.GoVersion)
		},
	}

	// 添加命令行工具
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(apiCmd)
	RootCmd.AddCommand(migrateCmd)

	return RootCmd
}

// Execute 启动命令
func Execute(app *App) {
	rootCmd := command(app)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 关闭客户端
	defer srv.Close()
}
