package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionPackageGroups() carapace.Action {
	return carapace.ActionExecCommand("pamac", "list", "--groups")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, line, "[group]")
		}
		return carapace.ActionValuesDescribed(vals...).Style(style.Yellow)
	}).Tag("groups")
}

func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) == 0 {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("pamac", "search", "--aur", c.CallbackValue)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ ]+) +(?P<installed>\[Installed\])? +(?P<version>[^ ]+) +(?P<repo>[^ ]+) $`)

			current := ""
			packages := make(map[string][]string)
			installed := make(map[string]bool)
			for _, line := range lines[:len(lines)-1] {
				if matches := r.FindStringSubmatch(line); matches != nil {
					current = matches[1]
					packages[current] = []string{fmt.Sprintf("[%v]", matches[4])}
					installed[current] = len(matches[2]) != 0
				} else {
					packages[current] = append(packages[current], strings.TrimSpace(line))
				}
			}

			vals := make([]string, 0)
			for key, value := range packages {
				s := style.Default
				if installed[key] {
					s = style.Blue
				}
				vals = append(vals, key, strings.Join(value, " "), s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func ActionInstalledPackages(explicit bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		arg := "-i"
		if explicit {
			arg = "-e"
		}
		return carapace.ActionExecCommand("pamac", "list", arg)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], fields[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})

	})
}
