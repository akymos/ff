package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	aliasCmd = &cobra.Command{
		Use:   "alias",
		Short: "Prints out the path to the alias file",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			err := internal.PopulateAlias()
			if err != nil {
				return err
			}
			fmt.Println(internal.GetAlias())
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(aliasCmd)
}
