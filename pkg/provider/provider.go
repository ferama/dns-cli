package provider

import "github.com/ferama/dns-cli/pkg/dnsrecord"

type Provider interface {
	// zone, recordType
	ListRecords(string, string) ([]dnsrecord.DnsRecord, error)
	AddRecord(string, dnsrecord.DnsRecord) error
	DeleteRecord(string, dnsrecord.DnsRecord) error
	UpdateRecord(string, old dnsrecord.DnsRecord, new dnsrecord.DnsRecord) error
	RefreshZone(string) error
}
