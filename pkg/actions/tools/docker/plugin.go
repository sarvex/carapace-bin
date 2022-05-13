package docker

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionPlugins completes plugins
//  TODO example
func ActionPlugins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "plugin", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
