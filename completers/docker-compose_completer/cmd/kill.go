package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill [OPTIONS] [SERVICE...]",
	Short: "Force stop service containers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	killCmd.Flags().StringP("signal", "s", "SIGKILL", "SIGNAL to send to the container.")
	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(killCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(killCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
