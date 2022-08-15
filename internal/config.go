package internal

import (
	"errors"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"os"
)

type config struct {
	AliasFile string
	DbFile    string
	Db        *bolt.DB
	FfDir     string
}

var BaseConfig config

func InitConfig() error {
	homedir, _ := os.UserHomeDir()
	BaseConfig = config{
		AliasFile: fmt.Sprintf("%s/.ff/ffAliases", homedir),
		DbFile:    fmt.Sprintf("%s/.ff/db", homedir),
		FfDir:     fmt.Sprintf("%s/.ff", homedir),
	}
	if _, err := os.Stat(BaseConfig.FfDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(BaseConfig.FfDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err := InitAlias()
	if err != nil {
		return err
	}
	db, err := InitDb()
	if err != nil {
		return err
	}
	BaseConfig.Db = db

	return nil
}
