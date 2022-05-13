package docker

import "github.com/rsteube/carapace"

// ActionDetachKeys completes detach keys
//   ctrl-@
//   ctrl-[
func ActionDetachKeys() carapace.Action {
	return carapace.ActionValues("{a-z}", `ctrl-\`, "ctrl-@", "ctrl-[", "ctrl-]", "ctrl-^", "ctrl-_", "ctrl-{a-z}")
}
