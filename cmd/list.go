package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
	"os"
)

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "short description",
		Long:    `long description`,
		Run: func(cmd *cobra.Command, args []string) {
			list := internal.LocalDb.FindAll()
			if len(list) == 0 {
				fmt.Println("Empty")
				os.Exit(0)
			}
			for k, v := range list {
				fmt.Printf("Alias: %s -> Folder: %s\n", k, v)
			}
			os.Exit(0)
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
