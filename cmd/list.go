package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "List saved aliases.",
		RunE: func(cmd *cobra.Command, args []string) error {
			list := internal.LocalDb.FindAll()
			if len(list) == 0 {
				fmt.Println("No aliases found.")
				return nil
			}
			for k, v := range list {
				fmt.Printf("Alias: %s -> Folder: %s\n", k, v)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
