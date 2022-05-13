package docker

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionSecrets completes secrets
//   another (updated 6 seconds ago)
//   example (updated 11 seconds ago)
func ActionSecrets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "secret", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Secret)
	})
}
