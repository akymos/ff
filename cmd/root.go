package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use: "ff",
		Long: `ff is a command-line tool to manage favorite folders, creating an alias,
to be used via shell directly with the cd command.`,
		Example: `$ ff add alias_name
$ ff add alias2 /tmp
$ ff update alias_name /var`,
	}
)

// Execute executes the root command.
func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Version = internal.GetVersion()
	rootCmd.SetVersionTemplate(fmt.Sprintf("Version: %s\nCommit: %s\nBuild on: %s\n", internal.GetVersion(), internal.GetCommit(), internal.GetDate()))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	return nil
}

func init() {
	err := internal.InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
