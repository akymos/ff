package cmd

import (
	"errors"
	"fmt"
	"github.com/creativeprojects/go-selfupdate"
	"github.com/spf13/cobra"
	"log"
	"os"
	"runtime"
)

var (
	selfUpdateCmd = &cobra.Command{
		Use:   "self-update",
		Short: "Update ff to the latest version.",
		Long:  `Update ff to the latest version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			latest, found, err := selfupdate.DetectLatest("akymos/ff")
			if err != nil {
				return fmt.Errorf("error occurred while detecting version: %v", err)
			}
			if !found {
				return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
			}

			if latest.LessOrEqual(rootCmd.Version) {
				log.Printf("Current version (%s) is the latest", rootCmd.Version)
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
		},
	}
)

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
