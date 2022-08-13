package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	aliasCmd = &cobra.Command{
		Use:   "alias",
		Short: "short description",
		Long:  `long description`,
		Run: func(cmd *cobra.Command, args []string) {
			err := internal.PopulateAlias()
			if err != nil {
				return
			}
			fmt.Println(internal.GetAlias())
		},
	}
)

func init() {
	rootCmd.AddCommand(aliasCmd)
}
