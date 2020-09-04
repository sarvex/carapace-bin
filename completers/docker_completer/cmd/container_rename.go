package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_renameCmd).Standalone()

	containerCmd.AddCommand(container_renameCmd)

	carapace.Gen(container_renameCmd).PositionalCompletion(
		action.ActionContainers(),
	)
}
