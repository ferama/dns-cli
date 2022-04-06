package ovhprovider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ferama/dns-cli/pkg/dnsrecord"
	"github.com/ferama/dns-cli/pkg/utils"
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
}

func NewOvhProvider() (*OvhProvider, error) {
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		return nil, err
	}

	p := OvhProvider{
		client: client,
	}

	return &p, nil
}

func (p *OvhProvider) getRecords(zone string, typeFilter string) ([]record, error) {
	var resp []int
	var records []record
	var err error
	if typeFilter != "" {
		typeFilter = strings.ToUpper(typeFilter)
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record?fieldType=%s", zone, typeFilter), &resp)
	} else {
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record", zone), &resp)
	}
	if err != nil {
		return nil, err
	}
	for _, recordId := range resp {
		var r record
		err = p.client.Get(fmt.Sprintf("/domain/zone/%s/record/%d", zone, recordId), &r)
		records = append(records, r)
	}
	return records, nil
}

func (p *OvhProvider) ListRecords(zone string, typeFilter string) ([]dnsrecord.DnsRecord, error) {
	var dnsrecords []dnsrecord.DnsRecord

	records, _ := p.getRecords(zone, typeFilter)
	for _, r := range records {
		dnsrecords = append(dnsrecords, dnsrecord.DnsRecord{
			Zone:      zone,
			Subdomain: r.SubDomain,
			Target:    r.Target,
			TTL:       int64(r.TTL),
			Type:      r.FieldType,
		})
	}
	return dnsrecords, nil
}

func (p *OvhProvider) AddRecord(zone string, record dnsrecord.DnsRecord) error {
	recordBody := recordFields{
		Target:    record.Target,
		TTL:       int(record.TTL),
		FieldType: record.Type,
		SubDomain: record.Subdomain,
	}
	fmt.Printf("adding an '%s' record for subdomain '%s' with target '%s'",
		recordBody.FieldType,
		recordBody.SubDomain,
		recordBody.Target)

	p.client.Post(fmt.Sprintf("/domain/zone/%s/record", zone), recordBody, nil)

	return nil
}

func (p *OvhProvider) DeleteRecord(zone string, record dnsrecord.DnsRecord) error {
	// p.client.Delete(fmt.Sprintf("/domain/zone/%s/record/%d", p.zone, record.ID), nil)
	all, _ := p.getRecords(zone, record.Type)
	ids := make([]int, 0)
	fmt.Println("I'm going to delete the following records")
	for _, r := range all {
		dnsr := dnsrecord.DnsRecord{
			Zone:      zone,
			Subdomain: r.SubDomain,
			Target:    r.Target,
			Type:      r.FieldType,
			TTL:       int64(r.TTL),
		}
		if dnsr.Match(record) {
			j, _ := json.Marshal(dnsr)
			fmt.Println(string(j))
			ids = append(ids, r.ID)
		}
	}
	if utils.AskForConfirmation("\nProceed?") {
		for _, id := range ids {
			p.client.Delete(fmt.Sprintf("/domain/zone/%s/record/%d", zone, id), nil)
		}
	}
	return nil
}

func (p *OvhProvider) RefreshZone(zone string) error {
	p.client.Post(fmt.Sprintf("/domain/zone/%s/refresh", zone), nil, nil)
	return nil
}
