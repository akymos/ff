package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "(Re)Generate the alias file",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			err := internal.GenerateAlias()
			if err != nil {
				return err
			}
			fmt.Println("Alias file (re)generated.")
			fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			internal.CheckNewVersion()
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
}
