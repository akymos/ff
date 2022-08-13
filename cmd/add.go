package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	addCmd = &cobra.Command{
		Use:     "add [alias] [path]",
		Aliases: []string{"a", "+"},
		Short:   "Create a directory alias",
		Long: `Create a directory alias.

Arguments:
[alias] required. The name of the alias.
[path] optional. It is the path to the directory to be aliased. Default is the current directory.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
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
			err := internal.LocalDb.Add(alias, path)
			if err != nil {
				return err
			}
			err = internal.PopulateAlias()
			if err != nil {
				return err
			}
			internal.WriteDb()
			fmt.Printf("Alias %s added for folder %s.\n", alias, path)
			fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
