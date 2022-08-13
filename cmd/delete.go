package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:     "delete [alias]",
		Aliases: []string{"del", "d", "-"},
		Short:   "Delete a directory alias.",
		Long: `Delete a directory alias.

Arguments:
[alias] required. The name of the alias.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			err := internal.LocalDb.Del(alias)
			if err != nil {
				return err
			}
			err = internal.PopulateAlias()
			if err != nil {
				return err
			}
			internal.WriteDb()
			fmt.Printf("Alias %s deleted.\n", alias)
			fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
