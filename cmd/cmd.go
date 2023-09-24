package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "fmarket",
		Short: "fmarket",
		Long:  "fmarket CLI",
		Run:   func(_ *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(CLIAppCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
