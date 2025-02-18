package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "workspace maintenance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workCmd).Standalone()

	rootCmd.AddCommand(workCmd)
}
