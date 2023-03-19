package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var installHooksCmd = &cobra.Command{
	Use:   "install-hooks",
	Short: "Install hook environments for all environments in the config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installHooksCmd).Standalone()

	installHooksCmd.Flags().String("color", "", "Whether to use color in output")
	installHooksCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	installHooksCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(installHooksCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config": carapace.ActionFiles(),
	})
}
