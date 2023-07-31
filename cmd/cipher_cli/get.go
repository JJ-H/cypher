package cmd

import (
	"cypher/internal/services"
	"cypher/internal/tools/crypto"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specified domain cypher, Usage: cypher get [domain]",
	Long:  `Get a specified domain cypher`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cypherAES := crypto.CypherAES
		domain := args[0]
		if domain == "" {
			color.Red("请输入 domain 以查询！")
			return
		}
		cypher := services.CredentialSrv.GetCredentialByDomain(domain)
		if cypher.Password == "" {
			color.Red("未查询到该域名的凭据！")
			return
		}
		decodedCypher, err := cypherAES.Decrypt(cypher.Password)
		if err != nil {
			color.Red("解密失败！")
			return
		}
		err = clipboard.WriteAll(decodedCypher)
		if err != nil {
			color.Red("复制密码到剪切板失败！")
			return
		}
		color.Green("密码已复制到剪切板！")
	},
}
