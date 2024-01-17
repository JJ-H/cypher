package database

import (
	"fmt"
	"os"
)

var (
	DB *os.File
)

func InitDB() *os.File {
	dir, err := os.UserHomeDir()
	var file *os.File
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(dir + "/cypher.json"); err != nil {
		// 自动创建文件并写入空内容
		file, err = os.Create(dir + "/cypher.json")
		fmt.Fprintf(file, "[]")
	}

	file, err = os.OpenFile(dir+"/cypher.json", os.O_RDWR|os.O_CREATE, 0644)
	os.Stat(dir + "/cypher.json")
	if err != nil {
		panic(err)
	}
	return file
}

func GetDB() *os.File {
	return InitDB()
}
