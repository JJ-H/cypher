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
	for _, cypher := range cyphers {
		if !plaintext {
			cypher.Password = "******"
		}
		table.Append([]string{cypher.Domain, cypher.Username, cypher.Password, cypher.Note})
	}
	table.Render()
}
