package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new tape file with example tape file contents and documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()
	rootCmd.AddCommand(newCmd)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
