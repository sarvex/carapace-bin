package docker

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionContainers completes container names
//   agitated_engelbart (alpine (Up 6 seconds))
//   crazy_satoshi (alpine (Up 4 seconds))
func ActionContainers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})", "--filter", "name="+c.CallbackValue)(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Style(styles.Docker.Container)
		})
	})
}

// ActionContainerIds completes container names
//   c84ca01b41f1 (alpine (Up 6 seconds))
//   1c3cf2aeee96 (alpine (Up 4 seconds))
func ActionContainerIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.ID}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
