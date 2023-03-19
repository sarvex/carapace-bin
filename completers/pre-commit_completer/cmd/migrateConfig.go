package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var migrateConfigCmd = &cobra.Command{
	Use:   "migrate-config",
	Short: "Migrate list configuration to new map configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateConfigCmd).Standalone()

	migrateConfigCmd.Flags().String("color", "", "Whether to use color in output")
	migrateConfigCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	migrateConfigCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(migrateConfigCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config": carapace.ActionFiles(),
	})
}
