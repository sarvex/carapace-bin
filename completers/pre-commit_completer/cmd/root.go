package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pre-commit",
	Short: "A framework for managing and maintaining multi-language pre-commit hooks",
	Long:  "https://github.com/pre-commit/pre-commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("version", "V", false, "show program's version number and exit")
}
