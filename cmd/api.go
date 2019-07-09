package cmd

import (
	"api/config"
	"api/log"
	"api/api"
	"strings"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:     "api",
	Aliases: []string{"server", "serve"},
	Short:   "Start HTTP API Server",
	Run: func(cmd *cobra.Command, args []string) {
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
		if addr != conf.Server.Addr {
			conf.Server.Addr = addr
		}
		apiCMD(conf)
	},
}

// API start http server
func apiCMD(conf *config.Config) {
	log.Debugf("配置: %+v", conf)
	if err := api.APIRun(conf); err != nil {
		log.Fatalf("run api error: %v", err)
		panic(err)
	}
}
