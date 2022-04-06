package provider

import "github.com/ferama/dns-cli/pkg/dnsrecord"

type Provider interface {
	// zone, recordType, subdomain like
	ListRecords(string, string, string) ([]dnsrecord.DnsRecord, error)
	AddRecord(string, dnsrecord.DnsRecord) error
	DeleteRecord(string, dnsrecord.DnsRecord) error
	RefreshZone(string) error
}
