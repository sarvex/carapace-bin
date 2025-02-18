package cmd

import (
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var work_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit go.work from tools or scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_editCmd).Standalone()
	work_editCmd.Flags().StringArrayS("dropreplace", "dropreplace", []string{}, "drop a replacement")
	work_editCmd.Flags().StringArrayS("dropuse", "dropuse", []string{}, "drop a use directive")
	work_editCmd.Flags().BoolS("fmt", "fmt", false, "reformat the go.work file without making other changes")
	work_editCmd.Flags().StringS("go", "go", "", "set the expected Go language version")
	work_editCmd.Flags().BoolS("json", "json", false, "print the final go.work in JSON format")
	work_editCmd.Flags().BoolS("print", "print", false, "print the final go.work in its text format")
	work_editCmd.Flags().StringArrayS("replace", "replace", []string{}, "add a replacement")
	work_editCmd.Flags().StringArrayS("use", "use", []string{}, "add a use directive")

	workCmd.AddCommand(work_editCmd)

	carapace.Gen(work_editCmd).FlagCompletion(carapace.ActionMap{
		"dropreplace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return golang.ActionWorkReplacements(c.Args[0])
			}
			return golang.ActionWorkReplacements("")
		}),
		"dropuse": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return golang.ActionWorkUses(c.Args[0])
			}
			return golang.ActionWorkUses("")
		}),
		"go": golang.ActionVersions(),
		"replace": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				if len(c.Args) > 0 {
					return golang.ActionWorkModules(c.Args[0]).Invoke(c).Suffix("=").ToA()
				}
				return golang.ActionWorkModules("").Invoke(c).Suffix("=").ToA()
			case 1:
				if util.HasPathPrefix(c.Value) {
					path, err := util.FindReverse(c.Dir, "go.work")
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}
					return carapace.ActionFiles().Chdir(filepath.Dir(path))
				}
				return golang.ActionModuleSearch()
			default:
				return carapace.ActionValues()
			}
		}),
		"use": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			path, err := util.FindReverse(c.Dir, "go.work")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionFiles().Chdir(filepath.Dir(path))

		}),
	})

	carapace.Gen(work_editCmd).PositionalCompletion(
		carapace.ActionFiles("go.work"),
	)
}
