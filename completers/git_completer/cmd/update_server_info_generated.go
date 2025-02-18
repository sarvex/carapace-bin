package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var update_server_infoCmd = &cobra.Command{
	Use:     "update-server-info",
	Short:   "Update auxiliary info file to help dumb servers",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(update_server_infoCmd).Standalone()
	update_server_infoCmd.Flags().BoolP("force", "f", false, "update the info files from scratch")
	rootCmd.AddCommand(update_server_infoCmd)
}
