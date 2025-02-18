package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:     "reload",
	Short:   "Reload one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()

	rootCmd.AddCommand(reloadCmd)

	carapace.Gen(reloadCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(reloadCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
