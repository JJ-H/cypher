package cmd

import (
	"cipher_manager/internal/services"
	"cipher_manager/internal/tools"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("plaintext", "P", false, "show plaintext password")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cyphers, Usage: cypher list [-P]",
	Long:  `List all cyphers`,
	Run: func(cmd *cobra.Command, args []string) {
		ciphers, err := services.CredentialSrv.ListCredential()
		if err != nil {
			fmt.Println("获取凭据列表失败!")
			return
		}
		if len(ciphers) == 0 {
			fmt.Println("暂无凭据!")
			return
		}

		plaintext, _ := cmd.Flags().GetBool("plaintext")

		tools.CypherPrinter(ciphers, plaintext)
	},
}
