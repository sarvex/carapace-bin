package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "Display a tree visualization of a dependency graph",
	Long:  "https://doc.rust-lang.org/cargo/commands/cargo-tree.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	rootCmd.Flags().Bool("all-features", false, "Activate all available features")
	rootCmd.Flags().String("charset", "", "Character set to use in output")
	rootCmd.Flags().String("color", "", "Coloring: auto, always, never")
	rootCmd.Flags().String("config", "", "Override a configuration value")
	rootCmd.Flags().String("depth", "", "Maximum display depth of the dependency tree")
	rootCmd.Flags().BoolP("duplicates", "d", false, "Show only dependencies which come in multiple versions")
	rootCmd.Flags().StringP("edges", "e", "", "The kinds of dependencies to display")
	rootCmd.Flags().String("exclude", "", "Exclude specific workspace members")
	rootCmd.Flags().StringP("features", "F", "", "Space or comma separated list of features to activate")
	rootCmd.Flags().StringP("format", "f", "", "Format string used for printing dependencies")
	rootCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().StringP("invert", "i", "", "Invert the tree direction and focus on the given package")
	rootCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	rootCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rootCmd.Flags().Bool("no-dedupe", false, "Do not de-duplicate (repeats all shared dependencies)")
	rootCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	rootCmd.Flags().Bool("offline", false, "Run without accessing the network")
	rootCmd.Flags().StringP("package", "p", "", "Package to be used as the root of the tree")
	rootCmd.Flags().String("prefix", "", "Change the prefix (indentation) of how each entry is displayed")
	rootCmd.Flags().String("prune", "", "Prune the given package from the display of the dependency tree")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.Flags().String("target", "", "Filter dependencies matching the given target-triple")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	rootCmd.Flags().Bool("workspace", false, "Display the tree for all packages in the workspace")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"charset": carapace.ActionValues("utf8", "ascii"),
	})
}
