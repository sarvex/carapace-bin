package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/precommit"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the pre-commit script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("allow-missing-config", false, "Whether to allow a missing `pre-commit` configuration file")
	installCmd.Flags().String("color", "", "Whether to use color in output")
	installCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	installCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	installCmd.Flags().StringP("hook-type", "t", "", "Hook type")
	installCmd.Flags().Bool("install-hooks", false, "Whether to install hook environments for all environments in the config file")
	installCmd.Flags().BoolP("overwrite", "f", false, "Overwrite existing hooks / remove migration mode")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"color":     carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config":    carapace.ActionFiles(),
		"hook-type": precommit.ActionHookTypes(),
	})
}
