package cmd

import (
	"fmt"

	ovhprovider "github.com/ferama/dns-cli/pkg/provider/ovh"
	"github.com/spf13/cobra"
)

func init() {
	domainCmd.AddCommand(listDomainCmd)
}

var listDomainCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := ovhprovider.NewOvhProvider()
		r, _ := provider.ListDomains()

		for _, item := range r {
			fmt.Println(item)
		}
	},
}
