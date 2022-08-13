package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	getCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
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
			p, err := internal.LocalDb.Get(alias)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%s", *p)
			os.Exit(0)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
