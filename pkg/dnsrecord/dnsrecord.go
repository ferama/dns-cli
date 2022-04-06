package dnsrecord

import "strings"

type DnsRecord struct {
	Zone string `json:"zone"`

	Subdomain string `json:"subdomain"`
	// The targets the DNS record points to
	Target string `json:"target"`
	// Type type of record, e.g. CNAME, A, SRV, TXT etc
	Type string `json:"recordType"`
	// TTL for the record
	TTL int64 `json:"recordTTL"`
}

func (d DnsRecord) Match(r DnsRecord) bool {
	if (r.Zone == "") || (r.Subdomain == "") {
		return false
	}
	if (r.Zone != d.Zone) || (r.Subdomain != d.Subdomain) {
		return false
	}
	if r.Target != "" && r.Target != d.Target {
		return false
	}
	if r.Type != "" && strings.ToUpper(r.Type) != strings.ToUpper(d.Type) {
		return false
	}
	return true
}
