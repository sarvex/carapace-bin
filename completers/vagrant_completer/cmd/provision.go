package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "provisions the vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(provisionCmd).Standalone()

	provisionCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	rootCmd.AddCommand(provisionCmd)

	carapace.Gen(provisionCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": action.ActionProvisioners().UniqueList(","),
	})

	carapace.Gen(provisionCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
