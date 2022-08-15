package internal

import (
	"errors"
	"fmt"
	"github.com/creativeprojects/go-selfupdate"
	"log"
	"os"
	"runtime"
)

var version string
var commit string
var date string

func CheckNewVersion() {
	latest, _, err := selfupdate.DetectLatest("akymos/ff")
	if err != nil {
		return
	}
	if latest.GreaterThan(version) {
		fmt.Printf("New version %s available.\nRun \"ff self-update\" to update.", version)
	}
}

func UpdateVersion() error {
	latest, found, err := selfupdate.DetectLatest("akymos/ff")
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %v", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}
	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}
	exe, err := os.Executable()
	if err != nil {
		return errors.New("could not locate executable path")
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %v", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}

func GetVersion() string {
	return version
}

func GetCommit() string {
	return commit
}

func GetDate() string {
	return date
}
