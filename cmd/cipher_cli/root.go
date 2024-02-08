package cmd

import (
	"cypher/internal/database"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cypher subCmd [options]",
	Short: "Cypher is a cli tool for cypher manager",
	Long: `You can easily manage your cyphers with cypher 
like set、get、list、delete and so on.`,
}

func init() {
	database.DB = database.InitDB()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
