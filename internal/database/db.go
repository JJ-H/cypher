package database

import (
	"fmt"
	"os"
)

var (
	DB *os.File
)

const (
	PersistFileName = "/.cypher.json"
)

func InitDB() *os.File {
	dir, err := os.UserHomeDir()
	var file *os.File
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(dir + PersistFileName); err != nil {
		// 自动创建文件并写入空内容
		file, err = os.Create(dir + PersistFileName)
		fmt.Fprintf(file, "[]")
	}

	file, err = os.OpenFile(dir+PersistFileName, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func GetDB() *os.File {
	return InitDB()
}
