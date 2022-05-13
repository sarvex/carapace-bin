package docker

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionRepositories completes repository names
//   alpine
//   bash
func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "images", "--format", "{{.Repository}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValues(vals[:len(vals)-1]...)
		})
	})
}

// ActionRepositoryTags completes repository names and tags separately
//   alpine:latest
//   bash:latest
func ActionRepositoryTags() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}", "--filter", "dangling=false")(func(output []byte) carapace.Action {
			// TODO add checks to not cause an out of bounds error
			lines := strings.Split(string(output), "\n")
			switch len(c.Parts) {
			case 0:
				reposWithSuffix := make([]string, len(lines)-1)
				for index, val := range lines[:len(lines)-1] {
					if val != " " {
						reposWithSuffix[index] = strings.SplitAfter(val, ":")[0]
					}
				}
				return carapace.ActionValues(reposWithSuffix...)
			case 1:
				tags := make([]string, 0)
				for _, val := range lines[:len(lines)-1] {
					if strings.HasPrefix(val, c.Parts[0]) {
						tag := strings.Split(val, ":")[1]
						tags = append(tags, tag)
					}
				}
				return carapace.ActionValues(tags...)
			default:
				return carapace.ActionValues()
			}
		}).Style(styles.Docker.Image)
	})
}
