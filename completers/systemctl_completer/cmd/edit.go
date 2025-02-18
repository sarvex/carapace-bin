package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit one or more unit files",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(editCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
