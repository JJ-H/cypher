package cmd

import (
	"cypher/internal/services/export"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("output", "o", "", "Export file path")
	exportCmd.Flags().StringP("type", "t", "csv", "Export file type, support json and csv")
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export cypher data to a file",
	Long:  `Export cypher data to a file，you can specify the file path with -o options and file type with -t options`,
	Run: func(cmd *cobra.Command, args []string) {
		output := cmd.Flag("output").Value.String()
		if output == "" {
			output = fmt.Sprintf("%s-%s.%s", "cypher", time.Now().Format("2006-01-02"), cmd.Flag("type").Value.String())
		}
		fileType := cmd.Flag("type").Value.String()
		var exporter export.Exporter
		switch fileType {
		case "csv":
			exporter = export.NewCsvExporter(",", output)
		default:
			color.Red("暂未支持的格式!")
			return
		}
		err := exporter.Export()
		if err != nil {
			color.Red("导出失败!")
			return
		}
		color.Green("导出成功，文件地址：%s ", output)
	},
}
