package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean out pre-commit files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().String("color", "", "Whether to use color in output")
	cleanCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})
}
