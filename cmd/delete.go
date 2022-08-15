package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type aliasList struct {
	Key string
	Val string
}

var (
	deleteCmd = &cobra.Command{
		Use:     "delete [alias]",
		Aliases: []string{"del", "d", "-"},
		Short:   "Delete a directory alias.",
		Long: `Delete a directory alias.

Arguments:
[alias] optional. The name of the alias.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			alias := ""
			if len(args) < 1 {
				list := internal.FindAll()
				if len(list) == 0 {
					return errors.New("no aliases found")
				}
				aliasesList := make([]aliasList, 0)
				for k, v := range list {
					aliasesList = append(aliasesList, aliasList{Key: k, Val: v})
				}
				aliasesList = append(aliasesList, aliasList{Key: "Exit", Val: "No delete any alias"})

				prompt := promptui.Select{
					Label: "Select an alias to delete:",
					Items: aliasesList,
					Templates: &promptui.SelectTemplates{
						Label:    "{{ . }}",
						Active:   "-> {{ .Key }} ({{ .Val }})",
						Inactive: "{{ .Key }} ({{ .Val }})",
						Selected: "-> {{ .Key }} ({{ .Val }})",
					},
				}

				i, _, err := prompt.Run()

				if err != nil {
					return err
				}

				if i == len(aliasesList)-1 {
					return nil
				}
				alias = aliasesList[i].Key
				err = internal.Del(alias)
				if err != nil {
					return err
				}
				err = internal.PopulateAlias()
				if err != nil {
					return err
				}
			} else {
				alias = args[0]
				err := internal.Del(alias)
				if err != nil {
					return err
				}
				err = internal.PopulateAlias()
				if err != nil {
					return err
				}
			}
			fmt.Printf("Alias %s deleted.\n", alias)
			fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
