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
	listRecordCmd.Flags().StringP("subdomain", "s", "", "filter by subdomain")
}

var listRecordCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")
		subdomain, _ := cmd.Flags().GetString("subdomain")

		typeFilter, _ := cmd.Flags().GetString("type-filter")

		provider, _ := ovhprovider.NewOvhProvider()
		r, _ := provider.ListRecords(zone, typeFilter, subdomain)

		w := tabwriter.NewWriter(os.Stdout, 5, 5, 5, ' ', 0)
		fmt.Fprintf(w, "Subdomain\tZone\tType\tTarget\n")
		header := "---------"
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", header, header, header, header)
		for _, item := range r {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", item.Subdomain, item.Zone, item.Type, item.Target)
		}
		w.Flush()
	},
}
