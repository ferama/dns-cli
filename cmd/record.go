package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(recordCmd)

	recordCmd.PersistentFlags().StringP("zone", "z", "", "dns zone")
	recordCmd.MarkPersistentFlagRequired("zone")
}

var recordCmd = &cobra.Command{
	Use:  "record",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}
