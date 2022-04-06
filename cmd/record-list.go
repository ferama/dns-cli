package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	ovhprovider "github.com/ferama/dns-cli/pkg/provider/ovh"
	"github.com/spf13/cobra"
)

func init() {
	recordCmd.AddCommand(listRecordCmd)

	listRecordCmd.Flags().StringP("type-filter", "t", "", "filter by record type")
}

var listRecordCmd = &cobra.Command{
	Use: "list [zone]",
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")

		typeFilter, _ := cmd.Flags().GetString("type-filter")
		provider, _ := ovhprovider.NewOvhProvider()
		r, _ := provider.ListRecords(zone, typeFilter)
		w := tabwriter.NewWriter(os.Stdout, 5, 5, 5, ' ', 0)
		for _, item := range r {
			if item.Subdomain != "" {
				fmt.Fprintln(w, fmt.Sprintf("%s.%s\t%s\t%s", item.Subdomain, item.Zone, item.Type, item.Target))
			} else {
				fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", item.Zone, item.Type, item.Target))
			}
		}
		w.Flush()
	},
}
