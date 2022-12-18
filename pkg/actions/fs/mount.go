package fs

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionMounts completes file system mounts
//
//	/boot/efi (/dev/sda1)
//	/dev (dev)
func ActionMounts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("/proc/mounts")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		lines := strings.Split(string(content), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, strings.Replace(fields[1], `\040`, " ", -1), fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	}).Tag("mounts")
}
