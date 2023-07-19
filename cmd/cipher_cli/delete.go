package cmd

import (
	"cipher_manager/internal/services"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a specified cypher, Usage: cypher delete [domain]",
	Long:  "Delete a specified domain cypher",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		if domain == "" {
			fmt.Println("请输入 domain 以删除！")
			return
		}
		services.CredentialSrv.DeleteCypherByDomain(domain)
		if services.CredentialSrv.GetCredentialByDomain(domain).Password == "" {
			fmt.Println("删除凭证成功！")
			return
		}
		fmt.Println("删除凭证失败！")
	},
}
