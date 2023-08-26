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
	listCmd.Flags().StringP("domain", "d", "", "query by domain")
	listCmd.Flags().StringP("note", "n", "", "query by note")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List cyphers, Usage: cypher list [options]",
	Long:  `List cyphers, support query by domain or note`,
	Run: func(cmd *cobra.Command, args []string) {
		domain := cmd.Flag("domain").Value.String()
		note := cmd.Flag("note").Value.String()

		var ciphers services.CredentialList
		var err error

		if domain != "" {
			ciphers, err = services.CredentialSrv.FuzzySearch(domain, "domain")
		} else if note != "" {
			ciphers, err = services.CredentialSrv.FuzzySearch(note, "note")
		} else {
			ciphers, err = services.CredentialSrv.ListCredential()
		}

		if err != nil {
			color.Red("获取凭据列表失败!")
			return
		}
		if len(ciphers) == 0 {
			color.Green("暂无凭据!")
			return
		}

		cypherAES := crypto.CypherAES
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
