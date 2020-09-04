package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_restartCmd).Standalone()

	container_restartCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container (default 10)")
	containerCmd.AddCommand(container_restartCmd)

	carapace.Gen(container_restartCmd).PositionalAnyCompletion(action.ActionContainers())
}
