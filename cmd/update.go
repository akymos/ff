package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	updateCmd = &cobra.Command{
		Use:     "update [alias] [path]",
		Aliases: []string{"upd", "u"},
		Short:   "Update a directory alias",
		Long: `Update a directory alias.

Arguments:
[alias] required. The name of the alias.
[path] optional. It is the path to the directory to be aliased. Default is the current directory.`,
		Example: `$ ff add alias_name /tmp
$ ff update alias_name /var`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			alias := args[0]
			path := ""
			if len(args) > 1 {
				path = args[1]
			}
			if path == "" || path == "." || path == "./" {
				p, err := os.Getwd()
				if err != nil {
					return err
				}
				path = p
			}
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				return err
			}
			alias = strings.ReplaceAll(alias, " ", "_")
			err := internal.Update(alias, path)
			if err != nil {
				return err
			}
			err = internal.GenerateAlias()
			if err != nil {
				return err
			}
			fmt.Printf("Alias %s updated to %s.\n", alias, path)
			fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			internal.CheckNewVersion()
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
