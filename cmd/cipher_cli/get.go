package cmd

import (
	"cipher_manager/internal/services"
	"fmt"
	"github.com/atotto/clipboard"
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
		domain := args[0]
		if domain == "" {
			fmt.Println("请输入 domain 以查询！")
			return
		}
		cypher := services.CredentialSrv.GetCredentialByDomain(domain)
		if cypher.Password == "" {
			fmt.Println("未查询到该域名的凭据！")
			return
		}
		clipboard.WriteAll(cypher.Password)
		fmt.Println("密码已复制到剪切板！")
	},
}
