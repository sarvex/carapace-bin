package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var validateConfigCmd = &cobra.Command{
	Use:   "validate-config",
	Short: "Validate .pre-commit-config.yaml files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateConfigCmd).Standalone()

	validateConfigCmd.Flags().String("color", "", "Whether to use color in output")
	validateConfigCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(validateConfigCmd)

	carapace.Gen(validateConfigCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(validateConfigCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
