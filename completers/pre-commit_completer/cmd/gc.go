package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Clean unused cached repos",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()

	gcCmd.Flags().String("color", "", "Whether to use color in output")
	gcCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(gcCmd)

	carapace.Gen(gcCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})
}
