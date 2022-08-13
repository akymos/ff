package internal

import (
	"errors"
	"fmt"
	"os"
)

type config struct {
	AliasFile string
	DbFile    string
	FfDir     string
}

var BaseConfig config

func InitConfig() error {
	homedir, _ := os.UserHomeDir()
	BaseConfig = config{
		AliasFile: fmt.Sprintf("%s/.ff/ffAlias", homedir),
		DbFile:    fmt.Sprintf("%s/.ff/db.json", homedir),
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
	InitDb()
	return nil
}
