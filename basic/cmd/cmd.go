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

func command(app *App) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   app.Name,
		Short: app.Short,
		Long:  app.Description,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			apiCMD(config.Conf)
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of " + app.Name,
		Long:  `All software has versions. This is ` + app.Name + `'s`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Name + " v" + app.Version + "\nGit Commit:" + app.GitCommit + "\nBuild Time:" + app.BuildTime + "\nGo Version:" + app.GoVersion)
		},
	}
	// 添加命令行工具
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(apiCmd)

	return rootCmd
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
