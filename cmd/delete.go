package cmd

import (
	"errors"
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"sort"
	"strings"
)

type aliasList struct {
	Status string
	Key    string
	Val    string
}

var (
	deleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm", "d", "-"},
		Short:   "Interactively allows you to delete an alias",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			alias := ""
			list := internal.FindAll()
			if len(list) == 0 {
				return errors.New("no aliases found")
			}
			aliasesList := make([]aliasList, 0)
			for k, v := range list {
				status := promptui.IconGood
				if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
					status = promptui.IconWarn
				}
				aliasesList = append(aliasesList, aliasList{
					Key:    k,
					Val:    v,
					Status: status,
				})
			}

			// sort aliasesList by key alphabetically
			sort.Slice(aliasesList, func(i, j int) bool {
				return aliasesList[i].Key < aliasesList[j].Key
			})

			_, height, err := terminal.GetSize(0)
			if err != nil {
				return err
			}

			prompt := promptui.Select{
				Label: "Select an alias to delete: (ctrl-c to exit)",
				Items: aliasesList,
				Size:  height - 3,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   fmt.Sprintf("%s {{ .Status }} {{ .Key }} ({{ .Val }})", promptui.IconSelect),
					Inactive: "{{ .Status }} {{ .Key }} ({{ .Val }})",
					Selected: "{{ .Status }} {{ .Key }} ({{ .Val }})",
				},
				Searcher: func(input string, index int) bool {
					alias := aliasesList[index]
					name := fmt.Sprintf("%s %s", alias.Key, alias.Val)
					input = strings.Replace(strings.ToLower(input), " ", "", -1)

					return strings.Contains(name, input)
				},
				Keys: &promptui.SelectKeys{
					Prev:     promptui.Key{Code: promptui.KeyPrev, Display: promptui.KeyPrevDisplay},
					Next:     promptui.Key{Code: promptui.KeyNext, Display: promptui.KeyNextDisplay},
					PageUp:   promptui.Key{Code: promptui.KeyBackward, Display: promptui.KeyBackwardDisplay},
					PageDown: promptui.Key{Code: promptui.KeyForward, Display: promptui.KeyForwardDisplay},
					Search:   promptui.Key{Code: readline.CharTab, Display: "TAB"},
				},
			}

			i, _, err := prompt.Run()

			if err != nil {
				if err == promptui.ErrInterrupt {
					return nil
				}
				return err
			}

			promptConfirm := promptui.Prompt{
				Label: "Are you sure to delete all saved data? (y/n)",
			}

			result, err := promptConfirm.Run()
			if err != nil {
				return err
			}

			if result == "y" {
				alias = aliasesList[i].Key
				err = internal.Del(alias)
				if err != nil {
					return err
				}
				err = internal.GenerateAlias()
				if err != nil {
					return err
				}
				fmt.Printf("Alias %s deleted.\n", alias)
				fmt.Printf("Now run \n%s\nor restart the shell.\n", "source \"$(ff alias)\"")
			}

			internal.CheckNewVersion()
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
