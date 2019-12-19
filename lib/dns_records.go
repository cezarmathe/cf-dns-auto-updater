package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	// DNSRecordType represents the dns record type that can be
	// updated
	DNSRecordType = "A"
)

// DNSRecord is a type that contains the body of a dns record that
// needs to be updated
type DNSRecord struct {
	Type       string      `json:"type"`
	Name       string      `json:"name"`
	Identifier *string     `json:"identifier,omitempty"`
	Content    *net.IPAddr `json:"content,omitempty"`
}

func (record *DNSRecord) String() string {
	return fmt.Sprintf("\n    Type: %s, \n    Name: %s, \n    Identifier: %s, \n    Content: %s\n",
		record.Type,
		record.Name,
		func() string {
			if record.Identifier == nil {
				return "none"
			}
			return *record.Identifier
		}(),
		func() string {
			if record.Content == nil {
				return "none"
			}
			return record.Content.String()
		}())
}

// DNSRecords is a type for multiple dns records
type DNSRecords []*DNSRecord

func (records DNSRecords) String() string {
	str := make([]string, len(records))

	for index, record := range records {
		str[index] = record.String()
	}

	return fmt.Sprintf("Records: %s", strings.Join(str, ""))
}

// NewDNSRecord creates a new update dns record
func NewDNSRecord(name string, ip *net.IPAddr) *DNSRecord {
	return &DNSRecord{
		Type:       DNSRecordType,
		Name:       name,
		Identifier: nil,
		Content:    ip,
	}
}

// JSON returns the json value of this UpdateDNSRecord
func (record *DNSRecord) JSON() (string, error) {
	out, err := json.Marshal(record)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Zone is a type that represents a domain name and a tld
// i.e. example.com
type Zone struct {
	Name       string     `json:"name"`
	Identifier *string    `json:"identifier"`
	Records    DNSRecords `json:"records"`
}

func (zone *Zone) String() string {
	return fmt.Sprintf("\n Zone - name: %s, identifier: %s, records: %s\n",
		(*zone).Name,
		func() string {
			if zone.Identifier == nil {
				return "none"
			}
			return *(*zone).Identifier
		}(),
		(*zone).Records)
}

// ProcessDNSList processes a list of domain names into zones and records
func ProcessDNSList(domainNameList []string) ([]*Zone, error) {
	// if we get an empty domain name list, we pass an error
	if len(domainNameList) == 0 {
		return nil, errors.New("empty domain name list")
	}

	// create a zone-name zone map for easier processing
	data := make(map[string]*Zone, 0)

	// process domain name list
	for _, domainName := range domainNameList {
		parts := strings.Split(domainName, ".")

		// return if a bad domain name was found
		if len(parts) < 2 {
			return nil, errors.New("domain name " + domainName + " is not valid")
		}

		// the zone name are the last two dot-separated words in the domain name
		zoneName := strings.Join(parts[len(parts)-2:], ".")

		// if the zone does not already exist, create it
		if data[zoneName] == nil {
			data[zoneName] = new(Zone)
		} else if data[zoneName].Name == "" {
			data[zoneName].Name = zoneName
		}

		// add the record to the zone
		zone := data[zoneName]
		dnsRecord := NewDNSRecord(domainName, nil)
		zone.Records = append(zone.Records, dnsRecord)
	}

	// convert zone map to zone slice
	zones := make([]*Zone, 0)

	for _, value := range data {
		zones = append(zones, value)
	}

	return zones, nil
}
