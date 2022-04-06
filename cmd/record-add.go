package cmd

import (
	"github.com/ferama/dns-cli/pkg/dnsrecord"
	ovhprovider "github.com/ferama/dns-cli/pkg/provider/ovh"
	"github.com/spf13/cobra"
)

func init() {
	recordCmd.AddCommand(addRecordCmd)

	addRecordCmd.Flags().StringP("type", "t", "", "record type")
	addRecordCmd.Flags().StringArrayP("targets", "i", make([]string, 0), "targts")
	addRecordCmd.MarkFlagRequired("type")
	addRecordCmd.Flags().StringP("subdomain", "s", "", "subdomain")

	addRecordCmd.MarkFlagRequired("targets")
	addRecordCmd.MarkFlagRequired("subdomain")
}

var addRecordCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")
		subdomain, _ := cmd.Flags().GetString("subdomain")

		rtype, _ := cmd.Flags().GetString("type")
		targets, _ := cmd.Flags().GetStringArray("targets")

		provider, _ := ovhprovider.NewOvhProvider()

		for _, t := range targets {
			r := dnsrecord.DnsRecord{
				Zone:      zone,
				Subdomain: subdomain,
				Type:      rtype,
				Target:    t,
				TTL:       0,
			}
			provider.AddRecord(zone, r)
		}
		provider.RefreshZone(zone)
	},
}
