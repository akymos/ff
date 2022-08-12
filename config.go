package main

import (
	"errors"
	"fmt"
	"os"
)

type config struct {
	aliasFile string
	dbFile    string
	ffDir     string
}

var baseConfig config

func initConfig() error {
	homedir, _ := os.UserHomeDir()
	baseConfig = config{
		aliasFile: fmt.Sprintf("%s/.ff/ffAlias", homedir),
		dbFile:    fmt.Sprintf("%s/.ff/db.json", homedir),
		ffDir:     fmt.Sprintf("%s/.ff", homedir),
	}
	if _, err := os.Stat(baseConfig.ffDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(baseConfig.ffDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err := initAlias()
	if err != nil {
		return err
	}
	initDb()
	return nil
}
