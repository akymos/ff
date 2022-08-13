package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	deleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "d", "-"},
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
			err := internal.LocalDb.Del(alias)
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
			fmt.Printf("Alias %s deleted\n", alias)
			fmt.Printf("Now run \n%s\nor restart the shell\n", "source \"$(ff alias)\"")
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
