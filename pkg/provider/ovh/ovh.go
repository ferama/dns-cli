package ovhprovider

import (
	"fmt"

	"github.com/ferama/dns-cli/pkg/dnsrecord"
	"github.com/ovh/go-ovh/ovh"
)

type recordFields struct {
	SubDomain string `json:"subDomain"`
	Target    string `json:"target"`
	FieldType string `json:"fieldType"`
	TTL       int    `json:"ttl"`
}

type record struct {
	recordFields
	Zone string `json:"zone"`
	ID   int    `json:"id"`
}

type OvhProvider struct {
	client *ovh.Client
	zone   string
}

func NewOvhProvider(zone string) (*OvhProvider, error) {
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		return nil, err
	}

	p := OvhProvider{
		client: client,
		zone:   zone,
	}

	// var resp []string
	// client.Get("/domain/zone", &resp)

	return &p, nil
}

func (p *OvhProvider) ListRecords(typeFilter string) ([]dnsrecord.DnsRecord, error) {
	var resp []int
	var records []dnsrecord.DnsRecord
	var err error
	if typeFilter != "" {
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record?fieldType=%s", p.zone, typeFilter), &resp)
	} else {
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record", p.zone), &resp)
	}
	if err != nil {
		return nil, err
	}
	for _, recordId := range resp {
		var r record
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record/%d", p.zone, recordId), &r)
		dnsName := p.zone
		if r.SubDomain != "" {
			dnsName = fmt.Sprintf("%s.%s", r.SubDomain, p.zone)
		}

		records = append(records, dnsrecord.DnsRecord{
			DNSName: dnsName,
			Target:  r.Target,
			TTL:     int64(r.TTL),
			Type:    r.FieldType,
		})
	}
	return records, nil
}
