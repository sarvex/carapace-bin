package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/precommit"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("all-files", "a", false, "Run on all the files in the repo")
	runCmd.Flags().String("checkout-type", "", "Indicates whether the checkout was a branch checkout")
	runCmd.Flags().String("color", "", "Whether to use color in output")
	runCmd.Flags().String("commit-msg-filename", "", "Filename to check when running during `commit-msg`")
	runCmd.Flags().String("commit-object-name", "", "Commit object name")
	runCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	runCmd.Flags().StringSlice("files", []string{}, "Specific filenames to run hooks on")
	runCmd.Flags().String("from-ref", "", "the original ref in a `from_ref...to_ref` diff expression")
	runCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	runCmd.Flags().String("hook-stage", "", "The stage during which the hook is fired")
	runCmd.Flags().String("is-squash-merge", "", "Indicates whether the merge was a squash merge")
	runCmd.Flags().String("local-branch", "", "Local branch ref used by `git push`")
	runCmd.Flags().StringP("origin", "o", "", "the destination ref in a `from_ref...to_ref` diff expression")
	runCmd.Flags().String("prepare-commit-message-source", "", "Source of the commit message")
	runCmd.Flags().String("remote-branch", "", "Remote branch ref used by `git push`")
	runCmd.Flags().String("remote-name", "", "Remote name used by `git push`")
	runCmd.Flags().String("remote-url", "", "Remote url used by `git push`")
	runCmd.Flags().String("rewrite-command", "", "Specifies the command that invoked the rewrite")
	runCmd.Flags().Bool("show-diff-on-failure", false, "When hooks fail, run `git diff` directly afterward")
	runCmd.Flags().StringP("source", "s", "", "the original ref in a `from_ref...to_ref` diff expression")
	runCmd.Flags().String("to-ref", "", "the destination ref in a `from_ref...to_ref` diff expression")
	runCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.AddCommand(runCmd)

	runCmd.Flag("files").Nargs = -1

	// TODO add missing flags
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"checkout-type": carapace.ActionValuesDescribed(
			"0", "file",
			"1", "branch",
		),
		"color":         carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"files":         carapace.ActionFiles(),
		"from-ref":      git.ActionRefs(git.RefOption{}.Default()),
		"hook-stage":    precommit.ActionHookStages(),
		"local-branch":  git.ActionLocalBranches(),
		"origin":        git.ActionRefs(git.RefOption{}.Default()),
		"remote-branch": git.ActionRemoteBranches(""),
		"source":        git.ActionRefs(git.RefOption{}.Default()),
		"to-ref":        git.ActionRefs(git.RefOption{}.Default()),
		"remote-name":   git.ActionRemotes(),
		"remote-url":    git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	})

	// TODO hook positional completion
}
