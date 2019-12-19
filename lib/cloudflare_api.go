package lib

import (
	"fmt"
	"net"

	"github.com/cloudflare/cloudflare-go"
)

var (
	// API is the cloudflare api
	API *cloudflare.API
)

// NewAPI initializes a new cloudflare api
func NewAPI(config *Config) error {
	// create a new cloudflare api
	api, err := cloudflare.New(config.AuthKey, config.AuthEmail)
	if err != nil {
		return err
	}

	// set the api object
	API = api

	return nil
}

// LoadIdentifiers check the zones from Cloudflare and updates the IPs if necessary
func LoadIdentifiers(zones *[]*Zone) error {

	// check that all zones exist and get their identifiers
	for _, zone := range *zones {
		id, err := API.ZoneIDByName((*zone).Name)
		if err != nil {
			return err
		}
		(*zone).Identifier = &id

		// check that all records exist and check their identifiers
		for _, dnsRecord := range (*zone).Records {
			// try to find matching records on cloudflare in this zone; only one should exist
			recs, err := API.DNSRecords(id, cloudflare.DNSRecord{
				Name: dnsRecord.Name,
				Type: dnsRecord.Type,
			})
			if err != nil {
				return err
			}

			// filter out bad responses
			if len(recs) == 0 {
				return fmt.Errorf("dns record %s does not exist on cloudflare", dnsRecord)
			}
			if len(recs) > 1 {
				return fmt.Errorf("more than one dns record on cloudflare exists for this specification %s", dnsRecord)
			}

			// get the identifier and the ip
			rec := recs[0]
			(*dnsRecord).Identifier = &(rec.ID)
			(*dnsRecord).Content, err = net.ResolveIPAddr("", rec.Content)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// PerformUpdate actually does the update
func PerformUpdate(ip *net.IPAddr, zones *[]*Zone) error {
	return nil
}
