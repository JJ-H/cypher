package tools

import (
	"cypher/internal/database/models"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CypherPrinter(cyphers models.CredentialList, plaintext bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{" 域名 ", " 用户名 ", " 密码 ", " 备注 "})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.BgRedColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.BgMagentaColor})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, cipher := range cyphers {
		if !plaintext {
			cipher.Password = "******"
		}
		table.Append([]string{cipher.Domain, cipher.Username, cipher.Password, cipher.Note})
	}
	table.Render()
}
