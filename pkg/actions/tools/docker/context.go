package docker

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionContexts completes context names
//   default (Current DOCKER_HOST based configuration)
//   example (custom context)
func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "context", "ls", "--format", "{{.Name}}\n{{.Description}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
