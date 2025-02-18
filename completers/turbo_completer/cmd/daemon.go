package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs the Turborepo background daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	rootCmd.AddCommand(daemonCmd)
}
