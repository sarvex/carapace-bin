package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var validateManifestCmd = &cobra.Command{
	Use:   "validate-manifest",
	Short: "Validate .pre-commit-hooks.yaml files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateManifestCmd).Standalone()

	validateManifestCmd.Flags().String("color", "", "Whether to use color in output")
	validateManifestCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(validateManifestCmd)

	carapace.Gen(validateManifestCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(validateManifestCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
