package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var realisation_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "query information about one or several realisations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(realisation_infoCmd).Standalone()

	realisation_infoCmd.Flags().Bool("json", false, "Produce output in JSON format")
	realisationCmd.AddCommand(realisation_infoCmd)

	addEvaluationFlags(realisation_infoCmd)
	addFlakeFlags(realisation_infoCmd)
	addLoggingFlags(realisation_infoCmd)

	// TODO positional completion
}
