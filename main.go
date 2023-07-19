package main

import (
	cmd "cipher_manager/cmd/cipher_cli"
	"cipher_manager/constant"
	"os"
)

func init() {
	constant.ProjectPath, _ = os.Getwd()
}

func main() {
	cmd.Execute()
}
