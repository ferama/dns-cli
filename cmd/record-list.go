package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	ovhprovider "github.com/ferama/dns-cli/pkg/provider/ovh"
	"github.com/spf13/cobra"
)

func init() {
	recordCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("type-filter", "t", "", "filter by record type")
}

var listCmd = &cobra.Command{
	Use:  "list [zone]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zone := args[0]

		typeFilter, _ := cmd.Flags().GetString("type-filter")
		provider, _ := ovhprovider.NewOvhProvider(zone)
		r, _ := provider.ListRecords(typeFilter)
		w := tabwriter.NewWriter(os.Stdout, 5, 5, 5, ' ', 0)
		for _, item := range r {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", item.DNSName, item.Type, item.Target))
		}
		w.Flush()
	},
}
