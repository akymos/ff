package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	addCmd = &cobra.Command{
		Use:     "add [alias] [path]",
		Aliases: []string{"a", "+"},
		Short:   "short description",
		Long:    `long description`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("you need, at least, specify the alias")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			alias := args[0]
			path := ""
			if len(args) > 1 {
				path = args[1]
			}
			if path == "" || path == "." || path == "./" {
				p, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
				}
				path = p
			}
			err := internal.LocalDb.Add(alias, path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = internal.PopulateAlias()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			internal.WriteDb()
			fmt.Printf("Alias %s added for folder %s\n", alias, path)
			fmt.Printf("Now run \n%s\nor restart the shell\n", "source \"$(ff alias)\"")
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
