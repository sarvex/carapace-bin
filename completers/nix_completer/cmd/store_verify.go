package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var store_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify the integrity of store paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_verifyCmd).Standalone()

	store_verifyCmd.Flags().Bool("no-contents", false, "Do not verify the contents of each store path")
	store_verifyCmd.Flags().Bool("no-trust", false, "Do not verify whether each store path is trusted")
	store_verifyCmd.Flags().StringP("sigs-needed", "n", "", "Require that each path is signed by at least n different keys")
	store_verifyCmd.Flags().StringP("substituter", "s", "", "Use signatures from the specified store")
	storeCmd.AddCommand(store_verifyCmd)

	addEvaluationFlags(store_verifyCmd)
	addFlakeFlags(store_verifyCmd)
	addLoggingFlags(store_verifyCmd)

	// TODO positional completion
}
