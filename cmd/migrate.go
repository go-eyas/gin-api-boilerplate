package cmd

import (
	"api/db"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		migrateCMD()
	},
}

func migrateCMD() {
	db.Migrate()
}
