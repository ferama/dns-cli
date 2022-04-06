package provider

import "github.com/ferama/dns-cli/pkg/dnsrecord"

type Provider interface {
	ListRecords(string) ([]dnsrecord.DnsRecord, error)
	AddRecord(dnsrecord.DnsRecord) error
	DeleteRecord(dnsrecord.DnsRecord) error
	UpdateRecord(old dnsrecord.DnsRecord, new dnsrecord.DnsRecord) error
}
