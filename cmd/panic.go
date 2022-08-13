package cmd

import (
	"fmt"
	"github.com/akymos/ff/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
)

var (
	panicCmd = &cobra.Command{
		Use:   "panic",
		Short: "WARINING!! This delete all saved data",
		RunE: func(cmd *cobra.Command, args []string) error {
			prompt := promptui.Prompt{
				Label: "Are you sure to delete all saved data? (y/n)",
			}

			result, err := prompt.Run()
			if err != nil {
				return err
			}

			if result == "y" {
				err := os.RemoveAll(internal.BaseConfig.FfDir)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Canceled")
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(panicCmd)
}
