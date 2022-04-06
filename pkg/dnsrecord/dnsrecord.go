package dnsrecord

type DnsRecord struct {
	DNSName string `json:"dnsName,omitempty"`
	// The targets the DNS record points to
	Target string `json:"target,omitempty"`
	// Type type of record, e.g. CNAME, A, SRV, TXT etc
	Type string `json:"recordType,omitempty"`
	// TTL for the record
	TTL int64 `json:"recordTTL,omitempty"`
	// // record id
	// ID string `json:"ID,omitempty"`
}
