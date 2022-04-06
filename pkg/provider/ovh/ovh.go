package ovhprovider

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
}

func NewOvhProvider() (*OvhProvider, error) {
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		return nil, err
	}

	p := OvhProvider{
		client: client,
	}

	// var resp []string
	// client.Get("/domain/zone", &resp)

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
	// register a record for each ip
	// for _, ip := range ips {
	// 	recordBody := recordFields{
	// 		Target:    ip,
	// 		TTL:       0,
	// 		FieldType: "A",
	// 		SubDomain: subdomain,
	// 	}
	// 	log.Printf("ADD A record for subdomain: %s, ip: %s", subdomain, ip)
	// 	p.client.Post(fmt.Sprintf("/domain/zone/%s/record", p.zone), recordBody, nil)
	// }

	// // set subdomain as owned
	// if !recordExists {
	// 	txtBody := recordFields{
	// 		Target:    p.getTxtOwner(),
	// 		TTL:       0,
	// 		FieldType: "TXT",
	// 		SubDomain: strings.TrimSuffix(host, "."+p.zone),
	// 	}
	// 	log.Printf("ADD TXT record for subdomain: %s", subdomain)
	// 	p.client.Post(fmt.Sprintf("/domain/zone/%s/record", p.zone), txtBody, nil)
	// }
	return nil
}

func (p *OvhProvider) DeleteRecord(zone string, record dnsrecord.DnsRecord) error {
	// p.client.Delete(fmt.Sprintf("/domain/zone/%s/record/%d", p.zone, record.ID), nil)
	all, _ := p.getRecords(zone, record.Type)
	for _, r := range all {
		dnsr := dnsrecord.DnsRecord{
			Zone:      zone,
			Subdomain: r.SubDomain,
			Target:    r.Target,
			Type:      r.FieldType,
			TTL:       int64(r.TTL),
		}
		if dnsr.Match(record) {
			// res = append(res, r)
			log.Printf("delete record with id %d", r.ID)
			j, _ := json.Marshal(dnsr)
			fmt.Println(string(j))
		}
	}
	return nil
}

func (p *OvhProvider) RefreshZone(zone string) error {
	p.client.Post(fmt.Sprintf("/domain/zone/%s/refresh", zone), nil, nil)
	return nil
}
