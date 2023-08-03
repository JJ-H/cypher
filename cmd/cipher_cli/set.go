package cmd

import (
	"cypher/internal/database/models"
	"cypher/internal/services"
	"cypher/internal/tools/crypto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"math/rand"
)

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("domain", "d", "", "cypher domain")
	setCmd.Flags().StringP("username", "u", "", "cypher username")
	setCmd.Flags().StringP("password", "p", "", "cypher password")
	setCmd.Flags().StringP("note", "n", "", "cypher note")
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a cypher for a specified domain，Usage: cypher set -d [domain] -u [username] -p [password]",
	Long:  `Set a cypher for a specified domain`,
	Run: func(cmd *cobra.Command, args []string) {
		cypherAES := crypto.CypherAES
		domain := cmd.Flag("domain").Value.String()
		username := cmd.Flag("username").Value.String()
		password := cmd.Flag("password").Value.String()
		note := cmd.Flag("note").Value.String()

		var err error
		if password, err = cypherAES.Encrypt(password); err != nil {
			color.Red("加密失败！")
			return
		}

		cypher := models.Credential{
			ID:       rand.Int63(),
			Domain:   domain,
			Username: username,
			Password: password,
			Note:     note,
		}
		_, err = services.CredentialSrv.SetCredential(cypher)
		if err == nil {
			color.Green("设置凭证成功！")
		}
	},
}
