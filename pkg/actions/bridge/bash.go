package bridge

import (
	"strings"

	_ "embed"

	"github.com/rsteube/carapace"
)

//go:embed bash.sh
var bash string

// ActionBash bridges completions registered in bash shell
func ActionBash(cmd string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(c.Args, c.CallbackValue)
		for index, arg := range args {
			arg = strings.Replace(arg, `\`, `\\`, -1)
			arg = strings.Replace(arg, `'`, `\'`, -1)
			arg = strings.Replace(arg, ` `, `\ `, -1)
			args[index] = arg
		}
		return carapace.ActionExecCommand("bash", "-c", bash, "--", strings.Join(args, " "))(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			nospace := false

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, "\t", 2)

				if len(splitted) > 1 {
					vals = append(vals, splitted...)
				} else {
					vals = append(vals, splitted[0], "")
				}
				if value := splitted[0]; !nospace && len(value) > 0 && strings.ContainsAny(value[len(value)-1:], `/=@:.,-`) {
					nospace = true
				}

			}
			a := carapace.ActionValuesDescribed(vals...)
			if nospace {
				a = a.NoSpace()
			}
			return a
		})
	})
}
