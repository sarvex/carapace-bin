package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/precommit"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the pre-commit script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().String("color", "", "Whether to use color in output")
	uninstallCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	uninstallCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	uninstallCmd.Flags().StringP("hook-type", "t", "", "Hook type")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"color":     carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config":    carapace.ActionFiles(),
		"hook-type": precommit.ActionHookTypes(),
	})
}
