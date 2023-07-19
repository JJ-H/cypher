package tools

import (
	"cipher_manager/internal/database/models"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CypherPrinter(cyphers models.CredentialList, plaintext bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{" 域名 ", " 用户名 ", " 密码 "})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiBlackColor, tablewriter.BgHiRedColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.BgGreenColor})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, cipher := range cyphers {
		if !plaintext {
			cipher.Password = "******"
		}
		table.Append([]string{cipher.Domain, cipher.Username, cipher.Password})
	}
	table.Render()
}
