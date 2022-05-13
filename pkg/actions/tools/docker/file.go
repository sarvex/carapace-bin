package docker

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionContainerPath completes container names and their file system separately
//   agitated_engelbart:/bin/echo
//   crazy_satoshi:/usr/lib/engines-1.1/afalg.so
func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionExecCommand("docker", "container", "ls", "--format", "{{.Names}}\n{{.Image}} ({{.Status}})")(func(output []byte) carapace.Action {
				vals := strings.Split(string(output), "\n")
				return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Invoke(c).Suffix(":").ToA().Style(styles.Docker.Container)
			})
		default:
			container := c.Parts[0]
			path := filepath.Dir(c.CallbackValue)

			args := []string{"exec", container, "ls", "-1", "-p", path}
			if splitted := strings.Split(c.CallbackValue, "/"); strings.HasPrefix(splitted[len(splitted)-1], ".") {
				args = append(args, "-a") // show hidden
			}
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return carapace.ActionExecCommand("docker", args...)(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")

					vals := make([]string, 0)
					for _, path := range lines[:len(lines)-1] {
						vals = append(vals, path, style.ForPathExt(path))
					}
					return carapace.ActionStyledValues(vals...)
				})
			})
		}
	})
}
