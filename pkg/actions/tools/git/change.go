package git

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
)

type ChangeOpts struct {
	Staged   bool
	Unstaged bool
}

func (o ChangeOpts) Default() ChangeOpts {
	o.Staged = true
	o.Unstaged = true
	return o
}

// ActionChanges completes (un)staged changes
//
//	fileA ( M)
//	pathA/fileB (??)
func ActionChanges(opts ChangeOpts) carapace.Action {
	// TODO multiparts action to complete step by step
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "status", "--porcelain")(func(output []byte) carapace.Action {
			if root, err := rootDir(c); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				untracked := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if len(line) > 3 {
						if (opts.Staged && line[1] == ' ') ||
							(opts.Unstaged && line[1] != ' ') {
							path := line[3:]
							if splitted := strings.SplitN(path, " -> ", 2); len(splitted) > 1 { // renamed
								path = splitted[1]
							}
							if relativePath, err := filepath.Rel(c.Dir, root+"/"+path); err != nil {
								return carapace.ActionMessage(err.Error())
							} else {
								if status := line[:2]; strings.Contains(status, "D") { // deleted
									untracked = append(untracked, relativePath, status, style.ForPathExt(relativePath, c))
								} else {
									untracked = append(untracked, relativePath, status, style.ForPath(relativePath, c))
								}
							}
						}
					}
				}
				return carapace.ActionStyledValuesDescribed(untracked...)
			}
		})
	})
}

// ActionRefChanges completes changes compared to given ref.
//
//	go.mod
//	cmd/carapace/main.go
func ActionRefChanges(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "diff-tree", "--name-only", "--no-commit-id", "-r", ref)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			path, err := util.FindReverse(c.Dir, ".git")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			relativeRoot, err := filepath.Rel(c.Dir, filepath.Dir(path))
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if relativeRoot == "." {
				relativeRoot = ""
			} else {
				relativeRoot += "/"
			}

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, relativeRoot+line)
			}

			return carapace.ActionValues(vals...).StyleF(style.ForPathExt)
		})
	})
}
