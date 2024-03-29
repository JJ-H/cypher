package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const VERSION = "1.3.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print cypher version, Usage: cypher version",
	Long:  `Print cypher version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cypher version ", VERSION)
	},
}
