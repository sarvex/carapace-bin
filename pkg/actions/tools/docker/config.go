package docker

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionConfigs completes config names
//   another (updated 4 seconds ago)
//   example (updated 23 seconds ago)
func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "config", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
