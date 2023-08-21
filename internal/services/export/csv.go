package export

import (
	"cypher/internal/services"
	"encoding/csv"
	"github.com/fatih/color"
	"os"
)

type CsvExporter struct {
	comma       string
	destination string
}

func NewCsvExporter(comma, destination string) *CsvExporter {
	if comma == "" {
		comma = ","
	}
	return &CsvExporter{comma: comma, destination: destination}
}

func (c CsvExporter) Export() error {
	srv := services.CredentialSrv
	var ciphers services.CredentialList
	var err error
	ciphers, err = srv.ListCredential()
	if err != nil {
		color.Red("获取凭据列表失败!")
		return err
	}
	if _, err = os.Stat(c.destination); err == nil {
		color.Red("当前路径已存在 %s !", c.destination)
		return nil
	}
	outFile, err := os.Create(c.destination)
	if err != nil {
		color.Red("创建文件失败!")
		return err
	}
	defer outFile.Close()
	_csv := csv.NewWriter(outFile)
	data := make([][]string, 0)
	data = append(data, []string{"Domain", "Username", "Password", "Note"})
	for _, v := range ciphers {
		data = append(data, []string{v.Domain, v.Username, v.Password, v.Note})
	}
	err = _csv.WriteAll(data)
	if err != nil {
		color.Red("写入文件失败!")
		return err
	}
	defer _csv.Flush()
	return nil
}
