package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

// TODO implement server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	Long:  `server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start server")
	},
}
