package cmd

import (
	"github.com/akymos/ff/internal"
	"github.com/spf13/cobra"
)

var (
	selfUpdateCmd = &cobra.Command{
		Use:   "self-update",
		Short: "Update ff to the latest version.",
		Long:  `Update ff to the latest version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			defer internal.BaseConfig.Db.Close()
			return internal.UpdateVersion()
		},
	}
)

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
