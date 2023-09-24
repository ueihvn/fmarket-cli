package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ueihvn/fmarket-cli/internal/cliapp"
)

func CLIAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cli",
		Short: "cli",
		Long:  "cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := cliapp.MustNewCliApp()
			if err := app.Start(); err != nil {
				return err
			}
			return nil
		},
	}
}
