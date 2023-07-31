package cmd

import (
	"cypher/internal/services"
	"cypher/internal/tools"
	"cypher/internal/tools/crypto"
	"github.com/fatih/color"
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
		cypherAES := crypto.CypherAES
		ciphers, err := services.CredentialSrv.ListCredential()
		if err != nil {
			color.Red("获取凭据列表失败!")
			return
		}
		if len(ciphers) == 0 {
			color.Green("暂无凭据!")
			return
		}

		for _, cypher := range ciphers {
			var password string
			if password, err = cypherAES.Decrypt(cypher.Password); err != nil {
				color.Red("解密失败！")
				return
			}
			cypher.Password = password
		}

		plaintext, _ := cmd.Flags().GetBool("plaintext")

		tools.CypherPrinter(ciphers, plaintext)
	},
}
