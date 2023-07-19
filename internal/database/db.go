package database

import (
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
	file, err = os.OpenFile(dir+"/cypher.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
