package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Help about any command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				vals := make([]string, 0)
				for _, cmd := range helpCmd.Root().Commands() {
					if !cmd.Hidden {
						vals = append(vals, cmd.Name(), cmd.Short)
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}).Tag("commands"),
			carapace.ActionValuesDescribed(
				"actions", "Learn about working with GitHub Actions",
				"environment", "Environment variables that can be used with gh",
				"exit-codes", "Exit codes used by gh",
				"formatting", "Formatting options for JSON data exported from gh",
				"mintty", "Information about using gh with MinTTY",
				"reference", "A comprehensive reference of all gh commands",
			).Style(style.Blue).Tag("help topics"),
		).ToA(),
	)
}
