package config

import (
	"cypher/internal/errors"
	"github.com/go-ini/ini"
	"os"
)

var (
	Config *ini.File
)

func InitConfig() error {
	var err error
	dir, err := os.UserHomeDir()
	if err != nil {
		return errors.LoadConfigError
	}

	Config, err = ini.Load(dir + "/.cypher_config.ini")
	if err != nil {
		return errors.LoadConfigError
	}
	return nil
}
