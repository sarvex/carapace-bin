package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/precommit"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var tryRepoCmd = &cobra.Command{
	Use:   "try-repo",
	Short: "Try the hooks in a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tryRepoCmd).Standalone()

	tryRepoCmd.Flags().BoolP("all-files", "a", false, "Run on all the files in the repo")
	tryRepoCmd.Flags().String("checkout-type", "", "Indicates whether the checkout was a branch checkout")
	tryRepoCmd.Flags().String("color", "", "Whether to use color in output")
	tryRepoCmd.Flags().String("commit-msg-filename", "", "Filename to check when running during `commit-msg`")
	tryRepoCmd.Flags().String("commit-object-name", "", "Commit object name")
	tryRepoCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	tryRepoCmd.Flags().String("files", "", "Specific filenames to run hooks on")
	tryRepoCmd.Flags().String("from-ref", "", "the original ref in a `from_ref...to_ref` diff expression")
	tryRepoCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	tryRepoCmd.Flags().String("hook-stage", "", "The stage during which the hook is fired")
	tryRepoCmd.Flags().String("is-squash-merge", "", "Indicates whether the merge was a squash merge")
	tryRepoCmd.Flags().String("local-branch", "", "Local branch ref used by `git push`")
	tryRepoCmd.Flags().StringP("origin", "o", "", "the destination ref in a `from_ref...to_ref` diff expression")
	tryRepoCmd.Flags().String("prepare-commit-message-source", "", "Source of the commit message")
	tryRepoCmd.Flags().String("remote-branch", "", "Remote branch ref used by `git push`")
	tryRepoCmd.Flags().String("remote-name", "", "Remote name used by `git push`")
	tryRepoCmd.Flags().String("remote-url", "", "Remote url used by `git push`")
	tryRepoCmd.Flags().String("rewrite-command", "", "Specifies the command that invoked the rewrite")
	tryRepoCmd.Flags().Bool("show-diff-on-failure", false, "When hooks fail, run `git diff` directly afterward")
	tryRepoCmd.Flags().StringP("source", "s", "", "the original ref in a `from_ref...to_ref` diff expression")
	tryRepoCmd.Flags().String("to-ref", "", "the destination ref in a `from_ref...to_ref` diff expression")
	tryRepoCmd.Flags().String("verbose,", "", "Verbose output")
	rootCmd.AddCommand(tryRepoCmd)

	carapace.Gen(tryRepoCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(tryRepoCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		// TODO hook completion
	)
}
