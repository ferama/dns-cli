package cmd

import (
	"github.com/ferama/dns-cli/pkg/dnsrecord"
	ovhprovider "github.com/ferama/dns-cli/pkg/provider/ovh"
	"github.com/spf13/cobra"
)

func init() {
	recordCmd.AddCommand(deleteRecordCmd)

	deleteRecordCmd.Flags().StringP("type", "t", "", "record type")
	deleteRecordCmd.Flags().StringP("subdomain", "s", "", "subdomain")

	deleteRecordCmd.MarkFlagRequired("subdomain")
}

var deleteRecordCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")
		subdomain, _ := cmd.Flags().GetString("subdomain")

		rtype, _ := cmd.Flags().GetString("type")

		r := dnsrecord.DnsRecord{
			Zone:      zone,
			Subdomain: subdomain,
			Type:      rtype,
		}

		provider, _ := ovhprovider.NewOvhProvider()
		provider.DeleteRecord(zone, r)
		provider.RefreshZone(zone)
	},
}
