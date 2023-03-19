package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/precommit"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var initTemplatedirCmd = &cobra.Command{
	Use:   "init-templatedir",
	Short: "Install hook script in a template directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initTemplatedirCmd).Standalone()

	initTemplatedirCmd.Flags().Bool("color", false, "Whether to use color in output")
	initTemplatedirCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	initTemplatedirCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	initTemplatedirCmd.Flags().StringP("hook-type", "t", "", "Hook type")
	initTemplatedirCmd.Flags().Bool("no-allow-missing-config", false, "Assume cloned repos should have a `pre-commit` config")
	rootCmd.AddCommand(initTemplatedirCmd)

	carapace.Gen(initTemplatedirCmd).FlagCompletion(carapace.ActionMap{
		"color":     carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config":    carapace.ActionFiles(),
		"hook-type": precommit.ActionHookTypes(),
	})

	carapace.Gen(initTemplatedirCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
