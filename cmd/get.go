package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:     "get [alias]",
		Aliases: []string{"g"},
		Short:   "Return the raw path of an alias.",
		Long: `Return the raw path of an alias, useful, for example, with ls.

Arguments:
[alias] required. The name of the alias.`,
		Example: `$ ls "$(ff get alias_name)"
$ cp test.txt "$(ff get alias_name)"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			alias := args[0]
			p, err := internal.Get(alias)
			if err != nil {
				return err
			}
			fmt.Printf("%s", *p)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
