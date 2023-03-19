package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var autoupdateCmd = &cobra.Command{
	Use:   "autoupdate",
	Short: "Auto-update pre-commit config to the latest repos versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoupdateCmd).Standalone()

	autoupdateCmd.Flags().Bool("bleeding-edge", false, "Update to the bleeding edge of `HEAD`")
	autoupdateCmd.Flags().String("color", "", "Whether to use color in output")
	autoupdateCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	autoupdateCmd.Flags().Bool("freeze", false, "Store \"frozen\" hashes in `rev` instead of tag names")
	autoupdateCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	autoupdateCmd.Flags().StringSlice("repo", []string{}, "Only update this repository")
	rootCmd.AddCommand(autoupdateCmd)

	carapace.Gen(autoupdateCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config": carapace.ActionFiles(),
		"repo":   git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	})
}
