package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func initQueryCmd(cmd *cobra.Command) {
	cmd.Flags().String("arch", "", "set an alternate architecture")
	cmd.Flags().String("cachedir", "", "set an alternate package cache location")
	cmd.Flags().BoolP("changelog", "c", false, "view the changelog of a package")
	cmd.Flags().CountP("check", "k", "check that package files exist (-kk for file properties)")
	cmd.Flags().String("color", "", "colorize the output")
	cmd.Flags().String("config", "", "set an alternate configuration file")
	cmd.Flags().Bool("confirm", false, "always ask for confirmation")
	cmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	cmd.Flags().Bool("debug", false, "display debug messages")
	cmd.Flags().BoolP("deps", "d", false, "list packages installed as dependencies [filter]")
	cmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	cmd.Flags().BoolP("explicit", "e", false, "list packages explicitly installed [filter]")
	cmd.Flags().StringP("file", "p", "", "query a package file instead of the database")
	cmd.Flags().BoolP("foreign", "m", false, "list installed packages not found in sync db(s) [filter]")
	cmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	cmd.Flags().BoolP("groups", "g", false, "view all members of a package group")
	cmd.Flags().String("hookdir", "", "set an alternate hook location")
	cmd.Flags().CountP("info", "i", "view package information (-ii for backup files)")
	cmd.Flags().BoolP("list", "l", false, "list the files owned by the queried package")
	cmd.Flags().String("logfile", "", "set an alternate log file")
	cmd.Flags().BoolP("native", "n", false, "list installed packages only found in sync db(s) [filter]")
	cmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	cmd.Flags().StringP("owns", "o", "", "query the package that owns <file>")
	cmd.Flags().BoolP("quiet", "q", false, "show less information for query and search")
	cmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	cmd.Flags().StringP("search", "s", "", "search locally-installed packages for matching strings")
	cmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	cmd.Flags().CountP("unrequired", "t", "list packages not (optionally) required by any package")
	cmd.Flags().BoolP("upgrades", "u", false, "list outdated packages [filter]")
	cmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arch":     carapace.ActionValues("i686", "x86_64"),
		"cachedir": carapace.ActionDirectories(),
		"color":    carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":   carapace.ActionFiles(),
		"dbpath":   carapace.ActionFiles(),
		"file":     carapace.ActionFiles(),
		"gpgdir":   carapace.ActionDirectories(),
		"hookdir":  carapace.ActionDirectories(),
		"logfile":  carapace.ActionFiles(),
		"owns":     carapace.ActionFiles(),
		"root":     carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if cmd.Flag("file").Changed {
				return carapace.ActionFiles()
			}
			return pacman.ActionPackages().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
