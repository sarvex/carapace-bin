package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var sampleConfigCmd = &cobra.Command{
	Use:   "sample-config",
	Short: "Produce a sample .pre-commit-config.yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sampleConfigCmd).Standalone()

	sampleConfigCmd.Flags().String("color", "", "Whether to use color in output")
	sampleConfigCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(sampleConfigCmd)

	carapace.Gen(sampleConfigCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})
}
