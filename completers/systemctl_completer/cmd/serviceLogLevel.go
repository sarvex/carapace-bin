package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

var serviceLogLevelCmd = &cobra.Command{
	Use:     "service-log-level",
	Short:   "Get/set logging threshold for service",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceLogLevelCmd).Standalone()

	rootCmd.AddCommand(serviceLogLevelCmd)

	carapace.Gen(serviceLogLevelCmd).PositionalCompletion(
		action.ActionServices(serviceLogLevelCmd),
		systemctl.ActionLogLevels(),
	)
}
