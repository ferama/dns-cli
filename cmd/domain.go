package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(domainCmd)

	// recordCmd.PersistentFlags().StringP("zone", "z", "", "dns zone")
	// recordCmd.MarkPersistentFlagRequired("zone")
}

var domainCmd = &cobra.Command{
	Use:  "domain",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}
