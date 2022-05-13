package docker

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionStacks completes stacks
//   // TODO example
func ActionStacks(orchestrator string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		arguments := []string{"stack", "ls", "--format", "{{.Name}}\n{{.Services}} on {{.Orchestrator}}"}
		if orchestrator == "swarm" || orchestrator == "kubernetes" {
			arguments = append(arguments, "--orchestrator", orchestrator)
		}
		return carapace.ActionExecCommand("docker", arguments...)(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		})
	})
}
