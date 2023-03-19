package precommit

import "github.com/rsteube/carapace"

// ActionHookTypes completes hook types
//
//	pre-commit
//	commit-msg
func ActionHookTypes() carapace.Action {
	return carapace.ActionValues(
		"pre-commit",
		"pre-merge-commit",
		"pre-push",
		"prepare-commit-msg",
		"commit-msg",
		"post-commit",
		"post-checkout",
		"post-merge",
		"post-rewrite",
	).Tag("hook types")
}

// ActionHookStages completes hook stages
//
//	commit
//	push
func ActionHookStages() carapace.Action {
	return carapace.ActionValues(
		"commit",
		"merge-commit",
		"prepare-commit-msg",
		"commit-msg",
		"post-commit",
		"manual",
		"post-checkout",
		"push",
		"post-merge",
		"post-rewrite",
	).Tag("hook stages")
}
